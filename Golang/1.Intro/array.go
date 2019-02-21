package main

import (
  "fmt"
)

func main() {
  // Method 1
  var age [10]int
  age[0] = 30
  fmt.Println(age)

  // Method 2
  height := [5]int{150, 160, 155, 190}
  fmt.Println(height)
}

// Output for method 1: [30 0 0 0 0 0 0 0 0 0]
// Output for method 2: [150 160 155 190 0]
