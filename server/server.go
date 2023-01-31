package server

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"example.com/echo/config"
)

func readCommand() {}

func respond() {}

func handleConnection(c net.Conn) {
	fmt.Println("Connection Success!")
}

func RunSyncServer() {

	var connection_count int = 0
	var addr = fmt.Sprintf("%s:%s", config.HOST, strconv.Itoa(config.PORT))

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		// handle error
		log.Fatalf("Failed to connect to the server: %s", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {

			// handle error
			log.Println("Failed to accept the connection")
		}
		connection_count += 1

		log.Println("Connected to a client, client count: ", connection_count)

		go handleConnection(conn)
	}
}
