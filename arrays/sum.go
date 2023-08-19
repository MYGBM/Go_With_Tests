package main

func Sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func SumAll(numbersToAdd ...[]int) []int {
	sum := []int{}
	for _, numbers := range numbersToAdd {
		sum = append(sum, Sum(numbers))

	}
	return sum
}
func SumAllTails(tailsToSum ...[]int) []int {
	sum := []int{}
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tails := numbers[1:]
			sum = append(sum, Sum(tails))
		}
	}
	return sum
}
