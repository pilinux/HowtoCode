package main

import (
	"sync"
	"time"
)

func study16() {
	// create a buffered channel of type int
	ch := make(chan int, 2) // capacity of 2
	println("Channel created successfully")

	var mu1, mu2 sync.Mutex
	mu1.Lock() // lock before starting goroutines
	mu2.Lock() // lock before starting goroutines

	// send 1 to channel ch in a separate goroutine
	go func() {
		defer mu1.Unlock()          // ensure the mutex is unlocked when the goroutine exits
		time.Sleep(5 * time.Second) // simulate some work
		ch <- 1
		println("Sent value to channel: 1")
	}()

	// send 2 to channel ch in a separate goroutine
	go func() {
		mu1.Lock()                  // wait for the first goroutine to finish
		defer mu2.Unlock()          // ensure the mutex is unlocked when the goroutine exits
		defer mu1.Unlock()          // unlock for good practice, not strictly needed here
		time.Sleep(5 * time.Second) // simulate some work
		ch <- 2
		println("Sent value to channel: 2")
	}()

	// send 3 to channel ch in a separate goroutine
	go func() {
		mu2.Lock()         // wait for the second goroutine to finish
		defer mu2.Unlock() // unlock for good practice, not strictly needed here
		ch <- 3
		println("Sent value to channel: 3")
	}()

	// receive 1st value from channel ch in the main goroutine
	value := <-ch

	// print the received value
	print("Received 1st value from channel: ")
	println(value)

	// receive 2nd value from channel ch in the main goroutine
	value = <-ch

	// print the received value
	print("Received 2nd value from channel: ")
	println(value)

	// receive 3rd value from channel ch in the main goroutine
	value = <-ch

	// print the received value
	print("Received 3rd value from channel: ")
	println(value)

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Sent value to channel: 3
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Received 3rd value from channel: 3
		Channel operations completed successfully

		Or,
		Channel created successfully
		Sent value to channel: 1
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Sent value to channel: 2
		Sent value to channel: 3
		Received 3rd value from channel: 3
		Channel operations completed successfully
	*/
}
