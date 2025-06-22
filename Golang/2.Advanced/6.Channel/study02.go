package main

func study02() {
	// create a channel of type int
	ch := make(chan int)
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

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		fatal error: all goroutines are asleep - deadlock!

		Reason:
		Both sending and receiving are happening in the same goroutine.
	*/
}
