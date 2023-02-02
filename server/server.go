package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"example.com/echo/config"
)

func readCommand(c net.Conn) (string, error) {

	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	log.Println(n, err)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(msg string, c net.Conn) error {
	if _, err := c.Write([]byte(msg)); err != nil {
		return err
	}

	return nil
}

func RunSyncServer() {
	var addr = fmt.Sprintf("%s:%s", config.HOST, strconv.Itoa(config.PORT))

	log.Println("Started a TCP server on", addr)

	var connection_count uint64 = 0

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("Listen failed: ", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept the connection", err)
		}

		connection_count += 1
		log.Println("Client connected with addr: ", conn.RemoteAddr(), "concurrent clients: ", connection_count)

		for {

			cmd, err := readCommand(conn)

			if err != nil {
				conn.Close()
				connection_count -= 1
				log.Println("Client disconnected: ", conn.RemoteAddr(), "concurrent clients: ", connection_count)

				if err == io.EOF {
					log.Println("Connection ended due to: ", err)
					break
				}

			}
			log.Println(cmd)
			err = respond(cmd, conn)
			if err != nil {
				log.Println("Failed to respond: ", err)
			}

		}

	}
}
