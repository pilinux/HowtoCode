package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting...")

	// Anonymous function
	go func() {
		fmt.Println("Executing at an unknown time.")
	}()
	time.Sleep(20 * time.Microsecond)

	fmt.Println("Another statement.")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Done.")
}

/*
** Output:
**
** Starting...
** Another statement.
** Executing at an unknown time.
** Done.
**/
