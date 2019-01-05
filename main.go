package main

import (
	"fmt"
	"github.com/marcelog/minesweeper-API/server"
	"os"
	"os/signal"
)

func main() {
	server := server.New(&server.Args{
		Address: "127.0.0.1",
		Port:    9999,
	})

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c)

	// Start the server
	fmt.Println("Hit ^C to quit")
	server.Run()

	// Block until any signal is received.
	s := <-c

	// Shutdown.
	fmt.Println("Got signal:", s)
	server.Stop()
}
