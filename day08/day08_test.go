package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccumulatorValue(t *testing.T) {
	actual := GetAccumulatorValue(data)
	assert.Equal(t, 5, actual)
}

func TestGetAccumulatorValueWithoutLoop(t *testing.T) {
	actual := GetAccumulatorValueWithoutLoop(data)
	assert.Equal(t, 8, actual)
}

var data = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}
