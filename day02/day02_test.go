package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateNumberOfValidPasswordsSledPolicy(t *testing.T) {
	actual := CalculateNumberOfValidPasswordsSledPolicy([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})

	assert.Equal(t, 2, actual)
}

func TestCalculateNumberOfValidPasswordsTobogganPolicy(t *testing.T) {
	actual := CalculateNumberOfValidPasswordsTobogganPolicy([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})

	assert.Equal(t, 1, actual)
}
