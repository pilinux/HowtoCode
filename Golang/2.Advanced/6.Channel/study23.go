package main

import (
	"fmt"
	"time"
)

// send-only function
func send(ch chan<- int, value int) {
	ch <- value
	fmt.Printf("Sent value to channel: %d\n", value)
}

// receive-only function
func recv(ch <-chan int, done chan<- bool) {
	v := <-ch
	fmt.Printf("Received value from channel: %d\n", v)
	time.Sleep(2 * time.Second) // simulate some processing time
	done <- true
}

func study23() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	fmt.Println("Channel created successfully")

	done := make(chan bool) // signal channel

	// send and receive 5 values using for loop in main func
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)

		go func() {
			send(ch, i)
		}()

		go func() {
			recv(ch, done)
		}()

		<-done // wait for receive to finish before next iteration
	}

	close(ch)
	fmt.Println("Channel closed after sending all values")
	fmt.Println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Iteration: 1
		Received value from channel: 1
		Iteration: 2
		Sent value to channel: 2
		Received value from channel: 2
		Iteration: 3
		Sent value to channel: 3
		Received value from channel: 3
		Iteration: 4
		Sent value to channel: 4
		Received value from channel: 4
		Iteration: 5
		Sent value to channel: 1
		Received value from channel: 5
		Sent value to channel: 5
		Channel closed after sending all values
		Channel operations completed successfully
	*/
}
