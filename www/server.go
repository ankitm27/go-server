package www

import socketServer "go-server/socket"

func RunPortServer() {
	go socketServer.CreateServer()
}
