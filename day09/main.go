package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day09/day09.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", FindInvalidNumber(values, 25))
	fmt.Printf("Solution part two: %d\n", FindEncryptionWeakness(values, 1504371145))
}
