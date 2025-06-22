package main

import (
	"fmt"
	"time"
)

func study20() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	fmt.Println("Channel created successfully")

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Sent value to channel: %d\n", i)
		}
		fmt.Println("Channel closed after sending all values")
	}()

	close(ch)

	for v := range ch {
		fmt.Printf("Received value from channel: %d\n", v)
	}

	fmt.Println("Channel operations completed successfully")

	time.Sleep(1 * time.Second) // Wait for goroutine to finish before exiting

	/*
		Output:
		Channel created successfully
		Channel operations completed successfully
		panic: send on closed channel
	*/
}
