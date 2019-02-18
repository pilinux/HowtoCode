package main

import (
  "fmt"
)

type Samurai struct {
  Name string
  Power int
}

func main() {
  roronoa := &Samurai {
    Name: "zoro",
    Power: 10000,
  }

  // Super changes value to the original roronoa
  Super(roronoa)

  fmt.Println(roronoa.Power)
}

func Super(s *Samurai) {
  s.Power += 20000
}

// Output: 30000
