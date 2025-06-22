package main

import (
	"fmt"
)

func study21() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	fmt.Println("Channel created successfully")

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Sent value to channel: %d\n", i)
		}
		close(ch)
		fmt.Println("Channel closed after sending all values")
	}()

	for v := range ch {
		fmt.Printf("Received value from channel: %d\n", v)
	}

	fmt.Printf("Received value from channel: %d\n", <-ch) // 0 since channel is closed
	fmt.Printf("Received value from channel: %d\n", <-ch) // 0 since channel is closed

	fmt.Println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received value from channel: 1
		Received value from channel: 2
		Received value from channel: 3
		Sent value to channel: 3
		Sent value to channel: 4
		Sent value to channel: 5
		Channel closed after sending all values
		Received value from channel: 4
		Received value from channel: 5
		Received value from channel: 0
		Received value from channel: 0
		Channel operations completed successfully
	*/
}
