package math

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(sets ...[]int) []int {
	var sums []int
	for _,set := range sets {
		sums = append(sums, Sum(set))
	}
	return sums
}

func SumAllTails(sets ...[]int) []int {
	var sums []int
	for _, set := range sets {
		if len(set) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(set[1:]))
		}
	}
	return sums
}