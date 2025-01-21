package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(array ...[]int) []int {
	var sums []int
	for _, numbers := range array {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(array ...[]int) []int {
	var sums []int
	for _, numbers := range array {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
