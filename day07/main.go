package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day07/day07.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateNumberOfBagsThatCanHoldAShinyGoldBag(values))
	fmt.Printf("Solution part two: %d\n", CalculateNumberOfIndividualBagsInsideTheShinyGoldBag(values))
}
