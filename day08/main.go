package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day08/day08.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", GetAccumulatorValue(values))
	fmt.Printf("Solution part two: %d\n", GetAccumulatorValueWithoutLoop(values))
}
