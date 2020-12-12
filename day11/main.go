package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day11/day11.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindNumberOfOccupiedSeatsPartOne(values))
	fmt.Printf("Solution part two: %d\n", FindNumberOfOccupiedSeatsPartTwo(values))
}
