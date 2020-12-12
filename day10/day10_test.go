package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheNumberOfJoltDifferences(t *testing.T) {
	actual := FindTheNumberOfJoltDifferences([]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4})
	assert.Equal(t, 35, actual)

	actual = FindTheNumberOfJoltDifferences([]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3})
	assert.Equal(t, 220, actual)
}

func TestFindTheNumberOfDistinctAdapterArrangements(t *testing.T) {
	actual := FindTheNumberOfDistinctAdapterArrangements([]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4})
	assert.Equal(t, 8, actual)

	actual = FindTheNumberOfDistinctAdapterArrangements([]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3})
	assert.Equal(t, 19208, actual)
}
