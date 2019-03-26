package variadic

import (
	"reflect"
	"testing"
)

func TestSumOfSlices(t *testing.T) {

	t.Run("Total 3 slices of numbers", func(t *testing.T) {
		numbersSlice1 := []int{1, 2, 3, 4, 5}
		numbersSlice2 := []int{6, 7, 8}
		numbersSlice3 := []int{6, 7, 8, 10}

		got := SumOfSlices(numbersSlice1, numbersSlice2, numbersSlice3)
		want := []int{15, 21, 31}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got -->%v<--, want -->%v<--", got, want)
		}
	})

}

// Note: invalid operation: got != want (slice can only be compared to nil)
// Explanation: a slice cannot be compared to a string
