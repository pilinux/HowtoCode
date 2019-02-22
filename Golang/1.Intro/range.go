package main

import (
  "fmt"
)

// Iterates over elements in a variety of data structures

func main() {
  // In a slice
  nums := []int{1, 22, 34, 47}
  for index, num := range nums {
    fmt.Println("index:", index, ", num:", num)
  }
  fmt.Println("")

  // In a map
  m := map[string]int{"k1": 10, "k2": 35}
  for key, value := range m {
    fmt.Printf("%s -->  %d\n", key, value)
  }
}

// Output:
/*
index: 0 , num: 1
index: 1 , num: 22
index: 2 , num: 34
index: 3 , num: 47

k1 -->  10
k2 -->  35
*/
