package main

func study01() {
	// create a channel of type int
	ch := make(chan int)
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
		fatal error: all goroutines are asleep - deadlock!

		Reason:
		Sending to a channel that has no receiver will block the goroutine indefinitely,
		leading to a deadlock.
	*/
}
