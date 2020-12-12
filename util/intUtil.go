package utils

import (
	"strconv"
)

func ToIntArray(values []string) (numbers []int) {
	for i := range values {
		n, _ := strconv.Atoi(values[i])
		numbers = append(numbers, n)
	}

	return numbers
}

func Max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func CreateAnArray(size, element int) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = element
	}
	return array
}
