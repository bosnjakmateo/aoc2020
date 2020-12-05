package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHighestSeatId(t *testing.T) {
	actual := CalculateHighestSeatId([]string{"FBFBBFFRLR"})
	assert.Equal(t, 357, actual)

	actual = CalculateHighestSeatId([]string{"BFFFBBFRRR"})
	assert.Equal(t, 567, actual)

	actual = CalculateHighestSeatId([]string{"FFFBBBFRRR"})
	assert.Equal(t, 119, actual)

	actual = CalculateHighestSeatId([]string{"BBFFBBFRLL"})
	assert.Equal(t, 820, actual)
}
