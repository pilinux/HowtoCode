// json.Unmarshal() method converts json(Byte data) into struct object

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonMsg := `{"Name":"Kate","Body":"Hello","Age":23}`

	type Response struct {
		Name string `json:"name"`
		Body string `json:"body"`
		Age  int    `json:age`
	}

	b := []byte(jsonMsg)
	fmt.Println(b)
	fmt.Println(string(b))

	// Create an empty Response struct
	var res Response
	// Unmarshal by passing a pointer to an empty struct
	json.Unmarshal(b, &res)
	fmt.Println(res)
	fmt.Println(res.Name)
}

// Output:
/*
[123 34 78 97 109 101 34 58 34 75 97 116 101 34 44 34 66 111 100 121 34 58 34 72 101 108 108 111 34 44 34 65 103 101 34 58 50 51 125]
{"Name":"Kate","Body":"Hello","Age":23}
{Kate Hello 23}
Kate
*/
