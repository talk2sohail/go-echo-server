package main

import (
	"flag"
	"log"

	"example.com/echo/config"
	"example.com/echo/server"
)

func setupFlags() {
	flag.StringVar(&config.HOST, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.PORT, "port", 7878, "port for the server")
	flag.Parse()
}
func main() {

	setupFlags()
	//fmt.Println("host no: ", config.HOST, " port number: ", config.PORT)
	log.Println("started the echo server at ", config.PORT)
	server.RunSyncServer()
}
