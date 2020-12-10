package main

import "sort"

func FindInvalidNumber(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		if !anyTwoNumbersHaveWantedSum(numbers[i-preamble:i], numbers[i]) {
			return numbers[i]
		}
	}

	return -1
}

func FindEncryptionWeakness(numbers []int, invalidNumber int) int {
	for i := 0; i < len(numbers); i++ {
		var summedNumbers []int
		sum := 0

		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			summedNumbers = append(summedNumbers, numbers[j])

			if sum == invalidNumber {
				sort.Ints(summedNumbers)
				return summedNumbers[0] + summedNumbers[len(summedNumbers)-1]
			}

			if sum > invalidNumber {
				break
			}
		}
	}

	return -1
}

func anyTwoNumbersHaveWantedSum(numbers []int, sum int) bool {
	for i, n1 := range numbers {
		for j, n2 := range numbers {
			if i != j && n1+n2 == sum {
				return true
			}
		}
	}

	return false
}
