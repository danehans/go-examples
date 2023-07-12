package main

import (
	"net"
	"log"
)

func main() {
	// Connect to the unix socket
	conn, err := net.Dial("unix", "/tmp/example.sock")
	if err != nil {
		log.Fatal("Error dialing unix socket: ", err)
	}
	defer conn.Close()

	// Send a message to the server
	_, err = conn.Write([]byte("Hello from client!"))
	if err != nil {
		log.Fatal("Error sending message: ", err)
	}
}
