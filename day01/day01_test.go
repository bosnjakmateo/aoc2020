package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateExpenseOfTwoEntries(t *testing.T) {
	actual := CalculateExpenseOfTwoEntries([]int{1721, 979, 366, 299, 675, 1456})

	assert.Equal(t, 514579, actual)
}

func TestCalculateExpenseOfThreeEntries(t *testing.T) {
	actual := CalculateExpenseOfThreeEntries([]int{1721, 979, 366, 299, 675, 1456})

	assert.Equal(t, 241861950, actual)
}
