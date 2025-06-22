package main

func study07() {
	// create a buffered channel of type int
	ch := make(chan int, 1) // capacity of 1
	println("Channel created successfully")

	// send 1 to channel ch in the main goroutine
	ch <- 1
	println("Sent value to channel: 1")

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
		fatal error: all goroutines are asleep - deadlock!

		Reason:
		Capacity of the channel is 1, so it can only hold one value at a time.
		Sending a second value without receiving the first one will cause a deadlock,
		as the channel is full and there is no receiver to take the value out.
	*/
}
