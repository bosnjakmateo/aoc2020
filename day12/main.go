package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day12/day12.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindManhattanDistance(values, false))
	fmt.Printf("Solution part two: %d\n", FindManhattanDistance(values, true))
}
