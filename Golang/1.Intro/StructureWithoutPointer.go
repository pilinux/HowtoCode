package main

import (
  "fmt"
)

// Create a structure
type Samurai struct {
  Name string
  Power int
}

func main() {
  // Create a variable `roronoa` using the structure
  roronoa := Samurai {
    Name: "zoro",
    Power: 10000,
  }

  // Super changes value to a copy of roronoa
  Super(roronoa)

  fmt.Println(roronoa.Power)
}

func Super(s Samurai) {
  s.Power += 20000
}

// Output: 10000
