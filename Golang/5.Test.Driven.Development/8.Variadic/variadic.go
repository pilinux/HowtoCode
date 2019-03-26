package variadic

// sum - add all the given numbers in a slice
func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumOfSlices - add all the given numbers in slices
func SumOfSlices(numbersToSum ...[]int) []int {
	var sums []int // empty slice

	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}

	return sums
}
