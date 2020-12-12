package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day10/day10.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindTheNumberOfJoltDifferences(values))
	fmt.Printf("Solution part two: %d\n", FindTheNumberOfDistinctAdapterArrangements(values))
}
