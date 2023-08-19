package main

func Sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func SumAll(numbersToAdd ...[]int) []int {
	length := len(numbersToAdd)
	sum := make([]int, length)
	for index, numbers := range numbersToAdd {
		sum[index] = Sum(numbers)

	}
	return sum
}
