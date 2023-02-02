package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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
	server.RunSyncServer()
	<-wait

}
