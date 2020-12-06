package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day06/day06.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateSumOfSameAnswersByAnyone(values))
	fmt.Printf("Solution part two: %d\n", CalculateSumOfSameAnswersByEveryone(values))
}
