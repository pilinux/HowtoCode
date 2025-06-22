package main

import "sync"

func study14() {
	// create a buffered channel of type int
	ch := make(chan int, 3) // capacity of 3
	println("Channel created successfully")

	// to control which value enters the channel first,
	// sync.WaitGroup to ensure the goroutine sending 1
	// runs before the one sending 2,
	// and the one sending 2 runs before the one sending 3
	var wg1, wg2 sync.WaitGroup
	wg1.Add(1) // we will wait for the 1st goroutine
	wg2.Add(1) // we will wait for the 2nd goroutine

	// send 1 to channel ch in a separate goroutine
	go func() {
		ch <- 1
		println("Sent value to channel: 1")
		wg1.Done() // signal that this goroutine is done
	}()

	// send 2 to channel ch in a separate goroutine
	go func() {
		wg1.Wait() // wait for the first goroutine to finish
		ch <- 2
		println("Sent value to channel: 2")
		wg2.Done() // signal that this goroutine is done
	}()

	// send 3 to channel ch in a separate goroutine
	go func() {
		wg2.Wait() // wait for the second goroutine to finish
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
		Received 1st value from channel: 1
		Sent value to channel: 2
		Sent value to channel: 3
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
