package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func FileInt(filePath string, separator string) (values []int) {
	return ToIntArray(FileString(filePath, separator))
}

func FileString(filePath string, separator string) (values []string) {
	bytes, _ := ioutil.ReadFile(filePath)

	str := string(bytes)

	return strings.Split(str, separator)
}

func ToIntArray(values []string) (numbers []int) {
	for i := range values {
		n, _ := strconv.Atoi(values[i])
		numbers = append(numbers, n)
	}

	return numbers
}

func Max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func CreateArray(size, element int) []int {
	array := make([]int, size)
	for i := range array {
		array[i] = element
	}
	return array
}

func Create2DStringArray(size int, defaultValue string) [][]string {
	array := make([][]string, size)
	for i := 0; i < size; i++ {
		innerArray := make([]string, size)
		for j := range innerArray {
			innerArray[j] = defaultValue
		}
		array[i] = innerArray
	}
	return array
}
