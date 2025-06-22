package main

import (
	"sync"
	"time"
)

func study18() {
	// create a buffered channel of type int
	ch := make(chan int, 2) // capacity of 2
	println("Channel created successfully")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	sentFirst := false

	// send 1 to channel ch in a separate goroutine
	go func() {
		mu.Lock()                   // lock to ensure only one goroutine sends at a time
		defer mu.Unlock()           // ensure the mutex is unlocked when done
		time.Sleep(5 * time.Second) // simulate some work
		ch <- 1
		println("Sent value to channel: 1")
		sentFirst = true
		cond.Signal()               // signal that the first value has been sent
		time.Sleep(5 * time.Second) // simulate some work
	}()

	// send 2 to channel ch in a separate goroutine
	go func() {
		mu.Lock()         // lock to ensure only one goroutine sends at a time
		defer mu.Unlock() // unlock for good practice, not strictly needed here
		// wait until the first value has been sent
		for !sentFirst {
			println("Waiting for the first value to be sent...")
			cond.Wait() // wait for the signal that the first value has been sent
		}
		ch <- 2
		println("Sent value to channel: 2")
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

	// (blocks forever)
	// receive 3rd value from channel ch in the main goroutine
	value = <-ch

	// print the received value
	print("Received 3rd value from channel: ")
	println(value)

	// never reaches here
	// receive 4th value from channel ch in the main goroutine
	value = <-ch

	// print the received value
	print("Received 4th value from channel: ")
	println(value)

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Waiting for the first value to be sent...
		Sent value to channel: 1
		Received 1st value from channel: 1
		Sent value to channel: 2
		Received 2nd value from channel: 2
		fatal error: all goroutines are asleep - deadlock!

		Note:
		- Only receive as many values as you send
	*/
}
