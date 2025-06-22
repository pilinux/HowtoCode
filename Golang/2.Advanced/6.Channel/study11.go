package main

func study11() {
	// create a buffered channel of type int
	ch := make(chan int, 2) // capacity of 2
	println("Channel created successfully")

	// send 1 to channel ch in a separate goroutine
	go func() {
		ch <- 1
		println("Sent value to channel: 1")
	}()

	// send 2 to channel ch in a separate goroutine
	go func() {
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
		Channel created successfully
		Sent value to channel: 2
		Received 1st value from channel: 2
		Sent value to channel: 1
		Received 2nd value from channel: 1
		Channel operations completed successfully

		Note: The order of "Sent value to channel" messages may vary due to goroutine scheduling.
	*/
}
