package main

func study05() {
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

	// send another value to the channel in a separate goroutine
	// caution: if not done in a separate goroutine, deadlock will occur
	go func() {
		ch <- 2
		println("Sent another value to channel: 2")
	}()

	// receive the second value from the channel
	secondValue := <-ch
	print("Received second value from channel: ")
	println(secondValue)

	// close the channel to signal that no more values will be sent
	close(ch)
	println("Channel operations completed successfully")

	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Sent another value to channel: 2
		Received second value from channel: 2
		Channel operations completed successfully
	*/
}
