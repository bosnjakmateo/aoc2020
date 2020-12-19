package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindEarliestBusIdMultipliedByWaitingTime(t *testing.T) {
	actual := FindEarliestBusIdMultipliedByWaitingTime(
		[]string{
			"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			"mem[8] = 11",
			"mem[7] = 101",
			"mem[8] = 0",
		},
	)
	assert.Equal(t, 165, actual)
}
