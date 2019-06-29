package socket

import (
	"fmt"
	utility "go-server/utility"
	"log"
	"net"
	"strconv"
)

var config = utility.GetConfig()
var socketPort = config.SocketPort

func CreateServer() {
	// fmt.Println("check11111111111111")
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(socketPort))

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", err)
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
