package practice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/hackerrank/practice"
)

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	assert.Equal(t, int32(15), practice.DiagonalDifference(arr))
}
