package main

import (
	utils "aoc2019/util"
	"sort"
)

func FindTheNumberOfJoltDifferences(adapters []int) int {
	adapters = sortAndAddValues(adapters, 0, utils.Max(adapters)+3)
	oneDiff, threeDiff := 0, 0

	for i := 0; i < len(adapters)-1; i++ {
		current := adapters[i]
		next := adapters[i+1]

		if next-current == 1 {
			oneDiff++
			continue
		}

		if next-current == 3 {
			threeDiff++
			continue
		}
	}

	return oneDiff * threeDiff
}

func FindTheNumberOfDistinctAdapterArrangements(adapters []int) (count int) {
	adapters = sortAndAddValues(adapters, 0, utils.Max(adapters)+3)
	combosUpToIndex := utils.CreateArray(len(adapters), 1)

	for i := 1; i < len(adapters); i++ {
		combosUpToIndex[i] = combosUpToIndex[i-1]

		if i > 1 && adapters[i]-adapters[i-2] <= 3 {
			combosUpToIndex[i] += combosUpToIndex[i-2]
		}
		if i > 2 && adapters[i]-adapters[i-3] <= 3 {
			combosUpToIndex[i] += combosUpToIndex[i-3]
		}
	}

	return combosUpToIndex[len(combosUpToIndex)-1]
}

func sortAndAddValues(numbers []int, min, max int) []int {
	numbers = append(numbers, min)
	sort.Ints(numbers)
	return append(numbers, max)
}
