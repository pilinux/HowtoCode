package main

import "fmt"

func study25() {
	ch := make(chan int, 1) // initialize a buffered channel with capacity 1
	fmt.Println("Channel created successfully")

	ch <- 1
	fmt.Println("Sent value to channel: 1")

	ch = nil  // simulate a nil channel
	close(ch) // runtime panic occurs here
	fmt.Println("Channel closed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		panic: close of nil channel

		nil channel:
		- closing a nil channel causes a runtime panic.
	*/
}
