package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	mask     string
	commands []Command
}

type Command struct {
	address, value int
}

func FindEarliestBusIdMultipliedByWaitingTime(data []string) (sum int) {
	operations := parseData(data)
	m := map[int]int{}

	for _, operation := range operations {
		for _, command := range operation.commands {
			brr := valueGoBrr(operation.mask, command.value)

			m[command.address] = brr
		}
	}

	for _, i := range m {
		sum += i
	}

	return sum
}

func valueGoBrr(mask string, value int) int {
	maskS := strings.Split(mask, "")
	valueString := strings.Split(fmt.Sprintf("%036s", strconv.FormatInt(int64(value), 2)), "")
	result := strings.Split(fmt.Sprintf("%036d", 0), "")

	for i := len(maskS) - 1; i >= 0; i-- {
		if mask[i] == 'X' {
			result[i] = valueString[i]
			continue
		}

		result[i] = maskS[i]
	}

	valueee := strings.Join(result, "")

	parseInt, _ := strconv.ParseInt(valueee, 2, 64)

	return int(parseInt)
}

func parseData(data []string) []*Operation {
	var operations []*Operation

	operation := &Operation{mask: data[0][7:len(data[0])]}

	for i := 1; i < len(data); i++ {
		if strings.Contains(data[i], "mask") {
			operations = append(operations, operation)
			operation = &Operation{mask: data[i][7:len(data[i])]}
			continue
		}

		match := addressAndValueMatch.FindAllStringSubmatch(data[i], 5)
		address, _ := strconv.Atoi(match[0][1])
		value, _ := strconv.Atoi(match[0][2])

		operation.commands = append(operation.commands, Command{address, value})
	}

	operations = append(operations, operation)

	return operations
}

var addressAndValueMatch = regexp.MustCompile(`mem\[(\d+).*=\s(\d+)`)
