// scan input stream line by line into a byte buffer (default behavior)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// To test Err()
	// os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		// Text() or Bytes() outputs only the last scanned line!!
		fmt.Println("Scanned text: ", in.Bytes())
		fmt.Println("Scanned text: ", in.Text())

	}

	if err := in.Err(); err != nil {
		fmt.Println("Error: ", err)
	}
}

// Either from terminal, give input
// or, feed a file:
// ./scan < scan.txt
