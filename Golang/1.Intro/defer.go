package main

import (
	"fmt"
)

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}

func fourth() {
	fmt.Println("Fourth")
}

func main() {
	defer second()
	first()
	defer third()
	fourth()
	second()
	defer first()

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output:
/*
First	// line 25
Fourth	// line 27
Second	// line 28
4		// for loop
3
2
1
0
First	// line 29
Third	// line 26
Second	// line 24
*/

// Example Usage:
/*
It is safe to close an open call, i.e `ReadWrite()`,
using deferred functions as they are executed in
LIFO order, (4 --> 3 --> 2 --> ...)
*/
