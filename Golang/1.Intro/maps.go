package main

import (
  "fmt"
)

func main() {
  // Create an empty map
  // make(map[keyType]valueType)
  m := make(map[string]int)
  fmt.Println("Print 1:", m)
  fmt.Println("Length of map:", len(m))
  fmt.Println("")

  m["key1"] = 10
  m["key2"] = 20
  fmt.Println("Print 2:", m)
  fmt.Println("Length of map:", len(m))
  fmt.Println("")

  // Assign values from map to variable
  var1, isKeyPresent := m["key1"]
  var2 := m["key2"]
  fmt.Println("var1:", var1)
  fmt.Println("var2:", var2)
  fmt.Println("Is the key present:", isKeyPresent)
  fmt.Println("")

  // Delete a key
  delete(m, "key1")
  _, isKeyPresent2 := m["key1"]
  fmt.Println("Print 3:", m)
  fmt.Println("Length of map:", len(m))
  fmt.Println("Is the key present:", isKeyPresent2)
  fmt.Println("")

}

// Output:
/*
Print 1: map[]
Length of map: 0

Print 2: map[key1:10 key2:20]
Length of map: 2

var1: 10
var2: 20
Is the key present: true

Print 3: map[key2:20]
Length of map: 1
Is the key present: false

*/
