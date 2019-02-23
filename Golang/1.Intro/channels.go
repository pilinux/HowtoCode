package main

import (
	"fmt"
)

func main() {
	// Make a channel
	message := make(chan string)

	go func() {
		// Send a value into the channel
		message <- "ping"
	}()
	msg := <-message
	fmt.Println(msg)
}

// Output:
/*
ping
*/
