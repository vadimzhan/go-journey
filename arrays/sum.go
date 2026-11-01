package main

func Sum(numbers [5]int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
