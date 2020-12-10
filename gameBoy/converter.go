package gameBoy

import "strconv"

func ConvertToInstruction(rawInstruction string) Instruction {
	value, _ := strconv.Atoi(rawInstruction[4:])

	return Instruction{
		Operation: rawInstruction[:3],
		value:     value,
		executed:  false,
	}
}
