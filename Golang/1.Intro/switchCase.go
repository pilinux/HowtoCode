package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	fmt.Println(today)
	fmt.Println(today.Day)

	t := today.Day()
	fmt.Println(t)

	switch t {
	case 1, 5, 10:
		fmt.Println("Go to gym.")
	case 15, 17, 22:
		fmt.Println("Buy some food.")
	case 24, 27, 31:
		fmt.Println("Party tonight!")
	case 2:
		fmt.Println("Watch a movie.")
	default:
		fmt.Println("Today is an interesting day!")
	}
}

// Output:
/*
2019-02-24 11:54:25.96305098 +0000 UTC m=+0.000375954
0x48e120
24
Party tonight!
*/
