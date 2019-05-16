package main

import (
	"fmt"
	"log"

	"github.com/firstrow/tcp_server"
)

func main() {
	server := tcp_server.New("localhost:9999")
	server.OnNewClient(func(c *tcp_server.Client) {
		c.Send("hello")
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		log.Println(message)
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		fmt.Println("client disconnect")
	})
	server.Listen()
}
