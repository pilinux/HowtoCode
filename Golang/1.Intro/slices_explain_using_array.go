package main

import (
  "fmt"
)

func main() {
  const array_capacity = 20
  array_length := 5

  var anArray [array_capacity]int

  for i := 0; i < array_length; i++ {
    anArray[i] = i * 2
  }
  fmt.Println(anArray)
  fmt.Print("Array capacity: ")
  fmt.Println(array_capacity)
  fmt.Print("Array length: ")
  fmt.Println(array_length)
  fmt.Println("")

  // Add a new item to the array
  anArray[array_length] = 11
  array_length++
  fmt.Println(anArray)
  fmt.Print("Array capacity: ")
  fmt.Println(array_capacity)
  fmt.Print("Array length: ")
  fmt.Println(array_length)
  fmt.Println("")
}

// Output:
/*
[0 2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Array capacity: 20
Array length: 5

[0 2 4 6 8 11 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Array capacity: 20
Array length: 6

*/
