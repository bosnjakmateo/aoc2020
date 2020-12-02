package main

import (
	"strconv"
	"strings"
)

func CalculateNumberOfValidPasswordsSledPolicy(passwordsData []string) (count int) {
	for _, passwordData := range passwordsData {
		min, max, letter, password := extractValidationData(passwordData)

		letterCount := strings.Count(password, letter)

		if letterCountIsWithinRange(letterCount, min, max) {
			count++
		}
	}

	return count
}

func CalculateNumberOfValidPasswordsTobogganPolicy(passwordsData []string) (count int) {
	for _, passwordData := range passwordsData {
		min, max, letter, password := extractValidationData(passwordData)

		firstLetter := string(password[min-1])
		secondLetter := string(password[max-1])

		if neededLetterAppearsOnce(letter, firstLetter, secondLetter) {
			count++
		}
	}

	return count
}

func extractValidationData(passwordData string) (int, int, string, string) {
	data := strings.Split(passwordData, " ")
	letterRange := strings.Split(data[0], "-")

	min, _ := strconv.Atoi(letterRange[0])
	max, _ := strconv.Atoi(letterRange[1])

	return min, max, data[1][:1], data[2]
}

func letterCountIsWithinRange(letterCount int, min int, max int) bool {
	return letterCount >= min && letterCount <= max
}

func neededLetterAppearsOnce(neededLetter string, letters ...string) bool {
	return strings.Count(strings.Join(letters, ""), neededLetter) == 1
}
