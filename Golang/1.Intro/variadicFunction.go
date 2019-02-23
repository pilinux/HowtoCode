package main

import (
  "fmt"
)

func sum(numbers ...int) {
  fmt.Print(numbers, " ")
  total := 0
  for _, number := range numbers {
    total += number
  }
  fmt.Println("=", total)
}

func main() {
  // Passing args
  sum(0)
  sum(1, 2)
  sum(3, 5, 8)

  // Passing args using array
  n := []int{1, 2, 3, 4, 5}
  sum(n ...)
}

// Output:
/*
[0] = 0
[1 2] = 3
[3 5 8] = 16
[1 2 3 4 5] = 15
*/
