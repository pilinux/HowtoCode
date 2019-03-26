package slice

import (
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Total 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got -->%d<--, want -->%d<--, given -->%v<--", got, want, numbers)
		}
	})

	t.Run("No limit on numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Got -->%d<--, want -->%d<--, given -->%v<--", got, want, numbers)
		}
	})

}
