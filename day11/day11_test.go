package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberOfOccupiedSeatsPartOne(t *testing.T) {
	actual := FindNumberOfOccupiedSeatsPartOne(data)
	assert.Equal(t, 37, actual)
}

func TestFindNumberOfOccupiedSeatsPartTwo(t *testing.T) {
	actual := FindNumberOfOccupiedSeatsPartTwo(data)
	assert.Equal(t, 26, actual)
}

var data = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}
