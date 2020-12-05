package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day04/day04.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateNumberOfValidPassports(values))
	fmt.Printf("Solution part two: %d\n", CalculateNumberOfValidPassportsV2(values))
}
