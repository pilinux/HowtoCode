/*
a pointer stores a memory address of a value

*/

package main

import "fmt"

func main() {
	var (
		counter int
		p       *int
		v       int
	)

	// zero value of a pointer is nil
	if p == nil {
		fmt.Printf("p is nil, it contains the address = %p\n", p)
		fmt.Println()
	}

	// assign the address of counter to p
	p = &counter
	if p == &counter {
		fmt.Printf("address saved in p = %p, address of counter in the memory = %p\n", p, &counter)
		fmt.Println()
	}

	// assign value to counter
	counter = 10
	fmt.Printf("address saved in p = %p, address of counter in the memory = %p\n", p, &counter)
	fmt.Printf("value saved in counter: using counter variable = %d, using pointer = %d\n", counter, *p)
	fmt.Printf("address of p in the memory where the pointer holds its data = %p\n", &p)
	fmt.Println()

	// v variable copies value from the location where p is pointing
	v = *p
	fmt.Printf("addess of v = %p, value of v = %d\n", &v, v)
	fmt.Println()

	fmt.Println("changing counter value...")
	counter = 20
	fmt.Printf("address saved in p = %p, address of counter in the memory = %p\n", p, &counter)
	fmt.Printf("value saved in counter: using counter variable = %d, using pointer = %d\n", counter, *p)
	fmt.Printf("address of p in the memory where the pointer holds its data = %p\n", &p)
	fmt.Printf("addess of v = %p, value of v = %d\n", &v, v)
	fmt.Println()

	// a function can change a value of another function or scope
	// through a pointer
	passPtrVal(&counter)
	fmt.Printf("value saved in counter: %d\n", counter)
}

func passPtrVal(pn *int) {
	// update value
	*pn++
}

// Output
/*
p is nil, it contains the address = 0x0

address saved in p = 0xc00012a000, address of counter in the memory = 0xc00012a000

address saved in p = 0xc00012a000, address of counter in the memory = 0xc00012a000
value saved in counter: using counter variable = 10, using pointer = 10
address of p in the memory where the pointer holds its data = 0xc000124018

addess of v = 0xc00012a008, value of v = 10

changing counter value...
address saved in p = 0xc00012a000, address of counter in the memory = 0xc00012a000
value saved in counter: using counter variable = 20, using pointer = 20
address of p in the memory where the pointer holds its data = 0xc000124018
addess of v = 0xc00012a008, value of v = 10

value saved in counter: 21
*/
