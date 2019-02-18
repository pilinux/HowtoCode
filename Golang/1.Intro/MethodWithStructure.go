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
  roronoa := &Samurai {
    Name: "zoro",
    Power: 11000,
  }

  roronoa.Super()

  fmt.Println(roronoa.Power)
}

// *Samurai is a receiver of Super method
func (s *Samurai) Super() {
  s.Power += 20000
}

// Output: 31000
