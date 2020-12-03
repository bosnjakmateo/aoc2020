package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day03/day03.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateNumberOfTreesForGivenSlope(values, Point{3, 1}))
	fmt.Printf("Solution part two: %d\n", CalculateProductForNumberOfTreesEncounteredForGivenSlopes(
		values,
		[]Point{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}),
	)
}
