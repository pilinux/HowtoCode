package main

func study08() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	println("Channel created successfully")

	// send 1 to channel ch in the main goroutine
	ch <- 1
	println("Sent value to channel: 1")

	// receive from channel ch in the main goroutine
	value := <-ch
	println("Received value from channel successfully")

	// print the received value
	print("Received value from channel: ")
	println(value)

	// send 2 to channel ch in the main goroutine
	ch <- 2
	println("Sent value to channel: 2")

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Sent value to channel: 2
		Channel operations completed successfully
	*/
}
