package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day13/day13.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindEarliestBusIdMultipliedByWaitingTime(values))
	fmt.Printf("Solution part two: %d\n", FindPerfectBusDepartures(values))
}
