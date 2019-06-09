package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func WsHandler() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8080")
	fmt.Println("ln", ln)
	conn, _ := ln.Accept()
	fmt.Println("conn", conn)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	newmessage := strings.ToUpper(message)
	conn.Write([]byte(newmessage + "\n"))
}

func main() {
	WsHandler()
}
