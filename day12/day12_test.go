package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindManhattanDistancePartOne(t *testing.T) {
	actual := FindManhattanDistance(
		[]string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		},
		false,
	)
	assert.Equal(t, 25, actual)
}

func TestFindManhattanDistancePartTwo(t *testing.T) {
	actual := FindManhattanDistance(
		[]string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		},
		true,
	)
	assert.Equal(t, 286, actual)
}
