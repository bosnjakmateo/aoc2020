package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindInvalidNumber(t *testing.T) {
	actual := FindInvalidNumber(data, 5)
	assert.Equal(t, 127, actual)
}

func TestFindEncryptionWeakness(t *testing.T) {
	actual := FindEncryptionWeakness(data, 127)
	assert.Equal(t, 62, actual)
}

var data = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}
