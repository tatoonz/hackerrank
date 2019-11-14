package practice

import (
	"math"
)

// DiagonalDifference calculate the absolute difference between the sums of its diagonals
// URL: https://www.hackerrank.com/challenges/diagonal-difference/problem
func DiagonalDifference(arr [][]int32) int32 {
	var (
		left  int32 = 0
		right int32 = 0
	)

	size := len(arr)
	for i := 0; i < size; i++ {
		left += arr[i][i]
		right += arr[i][size-i-1]
	}

	return int32(math.Abs(float64(left) - float64(right)))
}
