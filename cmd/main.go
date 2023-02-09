package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"example.com/echo/config"
	"example.com/echo/server"
)

type cleanupCallbackFunc func(chan struct{})

func setupFlags() {
	flag.StringVar(&config.HOST, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.PORT, "port", 7878, "port for the server")
	flag.Parse()
}

func shutdownAfter(t time.Duration) cleanupCallbackFunc {
	return func(wait chan struct{}) {
		// here you can do cleaning up thing..
		time.Sleep(t)
		fmt.Println("Gracefully shutting down the process.")
		close(wait)
	}
}

func waitForShutDown(cleanUp cleanupCallbackFunc) chan struct{} {
	wait := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c)

		s := <-c

		fmt.Println("Got signal: ", s)
		cleanUp(wait)
	}()

	return wait
}
func main() {
	setupFlags()

	wait := waitForShutDown(shutdownAfter(2 * time.Second))

	// NOTE: if you want to run multiple instances of the server, give 'addr' as ':0',
	// as this will be used to assign a random port number so that no server will be
	// listening to the same port i.e fixed port given by the user as flag argument
	addr := fmt.Sprintf("%s:%s", config.HOST, strconv.Itoa(config.PORT))
	s, err := server.NewServer(addr)
	if err != nil {
		log.Fatalf("Cannot spin up the server: %v", err)
	}

	s.Serve()

	<-wait

}
