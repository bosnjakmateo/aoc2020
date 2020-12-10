package main

import . "aoc2019/gameBoy"

func GetAccumulatorValue(rawInstructions []string) int {
	var instructions []Instruction

	for _, rawInstruction := range rawInstructions {
		instructions = append(instructions, ConvertToInstruction(rawInstruction))
	}

	gameBoy := GameBoy{Instructions: instructions}

	return gameBoy.Run()
}

func GetAccumulatorValueWithoutLoop(rawInstructions []string) (value int) {
	var instructions []Instruction
	var gameBoy = GameBoy{}

	for _, rawInstruction := range rawInstructions {
		instructions = append(instructions, ConvertToInstruction(rawInstruction))
	}

	for i := range instructions {
		originalOperation := instructions[i].Operation

		switch originalOperation {
		case "jmp":
			instructions[i].Operation = "nop"
		case "nop":
			instructions[i].Operation = "jmp"
		default:
			continue
		}

		gameBoy.Reset(append([]Instruction{}, instructions...))

		accumulatorValue, looped := gameBoy.RunWithoutLooping()

		if looped {
			instructions[i].Operation = originalOperation
		} else {
			value = accumulatorValue
			break
		}
	}

	return value
}
