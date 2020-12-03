package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNumberOfValidPasswordsSledPolicy(t *testing.T) {
	actual := CalculateNumberOfTreesForGivenSlope(
		landscape,
		Point{3, 1},
	)

	assert.Equal(t, 7, actual)
}

func TestCalculateProductForNumberOfTreesEncounteredForGivenSlopes(t *testing.T) {
	actual := CalculateProductForNumberOfTreesEncounteredForGivenSlopes(
		landscape,
		[]Point{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		},
	)

	assert.Equal(t, 336, actual)
}

var landscape = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}
