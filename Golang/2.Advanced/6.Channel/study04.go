package main

func study04() {
	// create a channel of type int
	ch := make(chan int)
	println("Channel created successfully")

	// send 1 to channel ch in a separate goroutine
	go func() {
		ch <- 1
		println("Sent value to channel: 1")
	}()

	// receive from channel ch in the main goroutine
	value := <-ch
	println("Received value from channel successfully")

	// print the received value
	print("Received value from channel: ")
	println(value)

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	// now, try to send another value to the channel
	// This will cause a panic because the channel is closed.
	ch <- 2

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Channel operations completed successfully
		panic: send on closed channel
	*/
}
