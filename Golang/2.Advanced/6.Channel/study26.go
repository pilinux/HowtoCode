package main

import (
	"fmt"
	"sync"
)

// split input channel into two output channels
func split(in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range in {
			select {
			case out1 <- v:
			case out2 <- v:
			}
		}
	}()
	return out1, out2
}

// merge two input channels into one output channel
func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range in1 {
			out <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range in2 {
			out <- v
		}
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func study26() {
	// Step 1: Input
	in := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	// Step 2: Split
	out1, out2 := split(in)

	// Process each split channel (for example, double the values)
	process := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for v := range in {
				out <- v * 2
			}
		}()
		return out
	}

	p1 := process(out1)
	p2 := process(out2)

	// Merge the processed channels
	merged := merge(p1, p2)

	// Print merged results
	fmt.Println("Merged results:")
	for v := range merged {
		fmt.Println(v)
	}

	/*
		Output:
		Merged results:
		4
		2
		8
		6
		10
		14
		16
		12
		18
		20

		Note:
		- The output may vary due to the non-deterministic nature of goroutines.
		- once you read from a channel (e.g., to print),
		those values are gone and cannot be read again by another goroutine or function.
	*/
}
