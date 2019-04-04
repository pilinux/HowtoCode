package main

import (
	"fmt"
	"reflect"
)

func main() {
	// `a` is a string
	a := "123"
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	// Convert string to byte
	b := stringToBytes(a)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)

	// Convert byte to string
	c := bytesToString(b)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
}

func bytesToString(data []byte) string {
	return string(data[:])
}

func stringToBytes(data string) []byte {
	return []byte(data)
}

// Output:
/*

string
123

[]uint8
[49 50 51]

string
123

**/
