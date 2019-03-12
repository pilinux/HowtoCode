// json.Marshal() method converts struct into Byte data

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Message struct {
		Name string
		Body string
		Age  int
	}

	m := Message{"Alice", "Hello", 20}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}

// Output:
/*
[123 34 78 97 109 101 34 58 34 65 108 105 99 101 34 44 34 66 111 100 121 34 58 34 72 101 108 108 111 34 44 34 65 103 101 34 58 50 48 125]
{"Name":"Alice","Body":"Hello","Age":20}
*/
