package main

import "sync"

func study13() {
	// create a buffered channel of type int
	ch := make(chan int, 2) // capacity of 2
	println("Channel created successfully")

	// to control which value enters the channel first,
	// sync.WaitGroup to ensure the goroutine sending 1
	// runs before the one sending 2
	var wg sync.WaitGroup
	wg.Add(1) // we will wait for one goroutine

	// send 1 to channel ch in a separate goroutine
	go func() {
		ch <- 1
		println("Sent value to channel: 1")
		wg.Done() // signal that this goroutine is done
	}()

	// send 2 to channel ch in a separate goroutine
	go func() {
		wg.Wait() // wait for the first goroutine to finish
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

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Sent value to channel: 1
		Sent value to channel: 2
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Channel operations completed successfully

		Explanation:
		- The first goroutine sends 1 and then calls wg.Done().
		- The second goroutine waits for wg.Done() (using wg.Wait()) before sending 2.
	*/
}
