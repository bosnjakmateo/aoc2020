package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day02/day02.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateNumberOfValidPasswordsSledPolicy(values))
	fmt.Printf("Solution part two: %d\n", CalculateNumberOfValidPasswordsTobogganPolicy(values))
}
