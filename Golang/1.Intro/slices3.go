package main

import (
  "fmt"
)

func main() {
  length := 0
  capacity := 3
  c := "Array Capacity: "
  l := "Slice Length: "

  scores := make([]int, length, capacity)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 100)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 200)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 300)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 400)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 500)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 600)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 700)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 800)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")

  scores = append(scores, 900)
  fmt.Print(c)
  fmt.Println(cap(scores))
  fmt.Print(l)
  fmt.Println(len(scores))
  fmt.Println(scores)
  fmt.Println("")
}

// Output:
/*
Array Capacity: 3
Slice Length: 0
[]

Array Capacity: 3
Slice Length: 1
[100]

Array Capacity: 3
Slice Length: 2
[100 200]

Array Capacity: 3
Slice Length: 3
[100 200 300]

Array Capacity: 6
Slice Length: 4
[100 200 300 400]

Array Capacity: 6
Slice Length: 5
[100 200 300 400 500]

Array Capacity: 6
Slice Length: 6
[100 200 300 400 500 600]

Array Capacity: 12
Slice Length: 7
[100 200 300 400 500 600 700]

Array Capacity: 12
Slice Length: 8
[100 200 300 400 500 600 700 800]

Array Capacity: 12
Slice Length: 9
[100 200 300 400 500 600 700 800 900]

*/
