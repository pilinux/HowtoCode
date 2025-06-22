package main

// Channels are Go’s way to communicate between goroutines (lightweight threads).
// They let you send and receive values safely between concurrent parts of your program.

func main() {
	// study01()
	// deadlock: unbuffered channel with no receiver
	/*
		Output:
		Channel created successfully
		fatal error: all goroutines are asleep - deadlock!

		Reason:
		Sending to a channel that has no receiver will block the goroutine indefinitely,
		leading to a deadlock.
	*/

	// study02()
	// deadlock: unbuffered channel with send and receive in the same goroutine
	/*
		Output:
		Channel created successfully
		fatal error: all goroutines are asleep - deadlock!

		Reason:
		Both sending and receiving are happening in the same goroutine.
	*/

	// study03()
	// send and receive in separate goroutines
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Channel operations completed successfully
	*/

	// study04()
	// panic: send on closed channel
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Channel operations completed successfully
		panic: send on closed channel
	*/

	// study05()
	// multiple sends and receives in separate goroutines
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

	// study06()
	// only send to a buffered channel
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Channel operations completed successfully

		Reason:
		Buffered channels allow sending values without a corresponding receiver immediately,
		so the send operation does not block, and the channel can be closed without causing a deadlock.
	*/

	// study07()
	// deadlock: buffered channel capacity full, not possible to send another value
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

	// study08()
	// buffered channel reused for multiple sends and receives
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Sent value to channel: 2
		Channel operations completed successfully
	*/

	// study09()
	// buffered channel reused for multiple sends and receives
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Received value from channel successfully
		Received value from channel: 1
		Sent value to channel: 2
		Received value from channel successfully
		Received value from channel: 2
		Channel operations completed successfully
	*/

	// study10()
	// buffered channel with more capacity
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Channel operations completed successfully
	*/

	// study11()
	// buffered channel: send multiple values in separate goroutines
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

	// study12()
	// buffered channel: send multiple values in one goroutine
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Channel operations completed successfully
	*/

	// study13()
	// buffered channel: synchronize two goroutines with WaitGroup
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

	// study14()
	// buffered channel: synchronize more than two goroutines with WaitGroup
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

	// study15()
	// buffered channel: synchronize two goroutines with mutex
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Channel operations completed successfully
	*/

	// study16()
	// buffered channel: synchronize more than two goroutines with mutex
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Sent value to channel: 3
		Received 1st value from channel: 1
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

	// study17()
	// buffered channel: synchronize two goroutines with condition variable
	/*
		Output:
		Channel created successfully
		Waiting for the first value to be sent...
		Sent value to channel: 1
		Sent value to channel: 2
		Received 1st value from channel: 1
		Received 2nd value from channel: 2
		Channel operations completed successfully

		Note:
		- cond.Wait() always unlocks the mutex while waiting
	*/

	// study18()
	// deadlock: tries to receive more values than sent
	/*
		Output:
		Channel created successfully
		Waiting for the first value to be sent...
		Sent value to channel: 1
		Received 1st value from channel: 1
		Sent value to channel: 2
		Received 2nd value from channel: 2
		fatal error: all goroutines are asleep - deadlock!

		Note:
		- Only receive as many values as you send
	*/

	// study19()
	// range over a channel
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received value from channel: 1
		Received value from channel: 2
		Received value from channel: 3
		Sent value to channel: 3
		Sent value to channel: 4
		Sent value to channel: 5
		Channel closed after sending all values
		Received value from channel: 4
		Received value from channel: 5
		Channel operations completed successfully

		Note: The order of "Sent value to channel" messages may vary due to goroutine scheduling.
	*/

	// study20()
	// panic: send on closed channel
	/*
		Output:
		Channel created successfully
		Channel operations completed successfully
		panic: send on closed channel
	*/

	// study21()
	// tries to receive values from a closed channel (will return zero value)
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received value from channel: 1
		Received value from channel: 2
		Received value from channel: 3
		Sent value to channel: 3
		Sent value to channel: 4
		Sent value to channel: 5
		Channel closed after sending all values
		Received value from channel: 4
		Received value from channel: 5
		Received value from channel: 0
		Received value from channel: 0
		Channel operations completed successfully
	*/

	// study22()
	// for select construct with channel
	/*
		Output:
		Channel created successfully
		Sent value to channel: 1
		Sent value to channel: 2
		Received value from channel: 1
		Received value from channel: 2
		Received value from channel: 3
		Sent value to channel: 3
		Sent value to channel: 4
		Sent value to channel: 5
		Received value from channel: 4
		Received value from channel: 5
		Channel closed after sending all values - 2
		Channel closed after sending all values - 1
		Channel operations completed successfully
	*/

	// study23()
	// send and receive functions for channels
	/*
		Output:
		Channel created successfully
		Iteration: 1
		Received value from channel: 1
		Iteration: 2
		Sent value to channel: 2
		Received value from channel: 2
		Iteration: 3
		Sent value to channel: 3
		Received value from channel: 3
		Iteration: 4
		Sent value to channel: 4
		Received value from channel: 4
		Iteration: 5
		Sent value to channel: 1
		Received value from channel: 5
		Sent value to channel: 5
		Channel closed after sending all values
		Channel operations completed successfully
	*/

	// study24()
	// nil channel
	/*
		Output:
		fatal error: all goroutines are asleep - deadlock!

		nil channel:
		- sending or receiving on a nil channel blocks forever.
		- closing a nil channel causes a runtime panic.
		- nil channels are often used to signal that no communication should happen.
	*/
}
