package main

func Sum(numbers []int) int {
	s := 0
	// range options is less performant than for i
	for _, n := range numbers {
		s += n
	}
	return s
}

// sum of all elements of array
func SumAll(numbersToSum ...[]int) (r []int) {
	for _, n := range numbersToSum {
		r = append(r, Sum(n))
	}
	return
}

// sum of every element of slices without head
func SumAllTails(numbers ...[]int) (r []int) {
	for _, n := range numbers {
		if len(n) == 0 {
			r = append(r, 0)
		} else {
			tails := n[1:] // think of this amazing stuff
			r = append(r, Sum(tails))
		}
	}
	return
}
