package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("Got -->%d<--, want -->%d<--, given -->%v<--", got, want, numbers)
	}
}

func BenchmarkArray(b *testing.B) {
	numbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleArray() {
	numbers := [5]int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}
