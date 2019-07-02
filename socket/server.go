package socket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-server/redis"
	"io"
	"log"
	"net"
	"strings"
)

// func CreateServer(port int) {
// 	fmt.Println("check")
// 	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

// 	if err != nil {
// 		log.Fatalf("Socket listen port %d failed,%s", port, err)
// 	}

// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			log.Fatalln(err)
// 			continue
// 		}
// 		go handleRequest(conn)
// 	}

// }

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
			// log.Println("Receive:", data)
			// redis.AddDataIntoRedis(data)
			isEOF := isTransportOver(data)
			// fmt.Println("is Eof", isEOF)
			if isEOF {
				data = strings.Replace(data, "\r\n\r\n", "", 1)
			}
			// fmt.Println("data1111", data)
			// fmt.Println("reflect", reflect.TypeOf(data))
			// var x interface{} = data
			// fmt.Println("x", x)
			// fmt.Println("reflect", reflect.TypeOf(x))
			// bytes, err := json.Marshal(data)
			// if err != nil {
			// 	fmt.Println("There is some problem, Please try again")
			// }
			// fmt.Println("bytes", bytes)
			// type dataObj struct {
			// 	data string
			// }
			// var data dataObj
			// err = json.Unmarshal(bytes, &data)
			// if err != nil {
			// 	fmt.Println("There is some problem, Please after some time",err)
			// }
			// data = `{"data":{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}}`
			rawIn := json.RawMessage(data)
			bytes, err := rawIn.MarshalJSON()
			if err != nil {
				fmt.Println("There is some problem,Please try after some time", err)
			}
			type dataObj struct {
				Data interface{}
				// LastName  string
				Key    string
				Secret string
			}
			var p dataObj
			err = json.Unmarshal(bytes, &p)
			if err != nil {
				fmt.Println("There is some problem, Please try after some time1", err)
			}
			// fmt.Println("bytes", string(bytes))
			// fmt.Println("p", p)
			// fmt.Println("data")
			// fmt.Println("data", p.Data)
			// fmt.Println("data", p.Key)
			// fmt.Println("data1111", p.Secret)
			redis.AddDataIntoRedis(data)
			if isEOF {
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
