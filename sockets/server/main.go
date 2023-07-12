package main

import (
	"net"
	"log"
	"io/ioutil"
)

func main() {
	// Listen on the unix socket
	listener, err := net.Listen("unix", "/tmp/example.sock")
	if err != nil {
		log.Fatal("Error listening on unix socket: ", err)
	}

	defer listener.Close()

	// Accept connection from client
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Error accepting connection: ", err)
	}
	defer conn.Close()

	// Read message from client
	message, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal("Error reading message: ", err)
	}

	log.Printf("Received message: %s\n", message)
}

