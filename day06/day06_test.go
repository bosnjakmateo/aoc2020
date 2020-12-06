package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateSumOfSameAnswersByAnyone(t *testing.T) {
	actual := CalculateSumOfSameAnswersByAnyone(forms)
	assert.Equal(t, 11, actual)
}

func TestCalculateSumOfSameAnswersByEveryone(t *testing.T) {
	actual := CalculateSumOfSameAnswersByEveryone(forms)
	assert.Equal(t, 6, actual)
}

var forms = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}
