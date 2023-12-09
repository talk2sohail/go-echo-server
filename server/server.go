package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var connection_count uint64 = 0

type Server struct {
	listener net.Listener
	quit     chan struct{}
	addr     string
}

func NewServer(addr string) (Server, error) {
	s := Server{
		quit: make(chan struct{}),
		addr: addr,
	}
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return Server{}, fmt.Errorf(fmt.Sprintf("Listen Error: %v", err))
	}

	s.listener = ln

	return s, nil
}

func readCommand(c net.Conn) (string, error) {

	var buf []byte = make([]byte, 512)
	fmt.Fprint(c, ">> ")

	n, err := c.Read(buf[:])
	log.Println(n, err)
	if err != nil {
		return "", err
	}

	userInput := string(buf[:n]) // this send the while user input
	msg := sanitize(userInput)
	cmd := match(msg)
	return cmd, nil
}

func sanitize(msg string) string {
	var out string
	out = msg[:len(msg)-1] // cut the new line
	out = strings.Trim(out, " ")
	return out
}

func match(cmd string) string {
	switch cmd {
	case "history":
		return "Great, you know your 'History'!"
	case "hello":
		return "Hello there, how are you doing?"
	default:
		return cmd
	}
}

func respond(msg string, c net.Conn) error {
	msg += "\n"
	if _, err := c.Write([]byte(msg)); err != nil {
		return err
	}

	return nil
}

func (s Server) Serve() {

	log.Println("Started a TCP server on", s.addr)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("Failed to accept the connection", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

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
