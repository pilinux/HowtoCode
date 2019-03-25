package integers

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	sum := add(3, 5)
	expected := 8

	if sum != expected {
		t.Errorf("Expected '%d', but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println("Sum is:", sum)
	// Output: Sum is: 6
}

/*
** The comment in the ExampleAdd function is a must to
** execute ExampleAdd function.
** Without that comment, the function will be compiled,
** but won't be executed.
** That comment acts as a 'want' while from the ExampleAdd
** function, we get 'got'
** got  '...'
** want '...'
 */
