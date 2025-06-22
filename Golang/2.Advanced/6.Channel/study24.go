package main

import "fmt"

func study24() {
	var ch chan int // nil channel
	go func() {
		ch <- 1   // blocks forever in this separate goroutine
		close(ch) // this line will never execute
		fmt.Println("Channel closed successfully")
	}()

	fmt.Println("Received value from channel:", <-ch) // blocks forever in main goroutine

	/*
		Output:
		fatal error: all goroutines are asleep - deadlock!

		nil channel:
		- sending or receiving on a nil channel blocks forever.
		- closing a nil channel causes a runtime panic.
		- nil channels are often used to signal that no communication should happen.
	*/
}
