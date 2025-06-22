package main

func study06() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	println("Channel created successfully")

	// send 1 to channel ch in the main goroutine
	ch <- 1
	println("Sent value to channel: 1")

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Channel operations completed successfully

		Reason:
		Buffered channels allow sending values without a corresponding receiver immediately,
		so the send operation does not block, and the channel can be closed without causing a deadlock.
	*/
}
