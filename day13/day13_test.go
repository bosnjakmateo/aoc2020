package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindEarliestBusIdMultipliedByWaitingTime(t *testing.T) {
	actual := FindEarliestBusIdMultipliedByWaitingTime(
		[]string{
			"939",
			"7,13,x,x,59,x,31,19",
		},
	)
	assert.Equal(t, 295, actual)
}

func TestFindPerfectBusDepartures(t *testing.T) {
	actual := FindPerfectBusDepartures(
		[]string{
			"939",
			"17,x,13,19",
		},
	)
	assert.Equal(t, 3417, actual)

	actual = FindPerfectBusDepartures(
		[]string{
			"939",
			"7,13,x,x,59,x,31,19",
		},
	)
	assert.Equal(t, 1068781, actual)

	actual = FindPerfectBusDepartures(
		[]string{
			"939",
			"1789,37,47,1889",
		},
	)
	assert.Equal(t, 1202161486, actual)
}
