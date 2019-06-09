package socket

import (
	"bufio"
	redis "go-server/redis"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func CreateServer(port int) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		// w   = bufio.NewWriter(conn)
	)
ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])
		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			redis.AddDataIntoRedis(data)
			if isTransportOver(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}
	}
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return over
}
