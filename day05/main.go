package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day05/day05.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateHighestSeatId(values))
	fmt.Printf("Solution part two: %d\n", CalculateYourSeatId(values))
}
