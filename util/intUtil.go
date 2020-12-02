package utils

import "strconv"

func ToIntArray(values []string) (numbers []int) {
	for i := range values {
		n, _ := strconv.Atoi(values[i])
		numbers = append(numbers, n)
	}

	return numbers
}
