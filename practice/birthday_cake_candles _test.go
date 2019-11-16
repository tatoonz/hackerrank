package practice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tatoonz/hackerrank/practice"
)

func TestBirthdayCakeCandles(t *testing.T) {
	assert.Equal(t, int32(2), practice.BirthdayCakeCandles([]int32{3, 2, 1, 3}))
}
