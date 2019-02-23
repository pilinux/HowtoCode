package main

import (
	"fmt"
)

func main() {
	// Make a channel of strings buffering up to 2 values
	message := make(chan string, 2)

	go func() {
		// Send 2 values into the channel
		message <- "ping"
		message <- "pong"
	}()
	msg := <-message
	fmt.Println(msg)
	msg = <-message
	fmt.Println(msg)
}

// Output for method 1:
/*
ping
pong
*/
