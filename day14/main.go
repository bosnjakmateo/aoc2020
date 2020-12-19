package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day14/day14.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindEarliestBusIdMultipliedByWaitingTime(values))
}
