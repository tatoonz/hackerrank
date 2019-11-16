package practice

import (
	"fmt"
	"math"
)

// MiniMaxSum prints the minimum and maximum values that can be calculated by summing exactly four of the five integers
// URL: https://www.hackerrank.com/challenges/mini-max-sum/problem
func MiniMaxSum(arr []int32) {
	total := 0
	min := math.MaxInt64
	max := 0
	for _, v := range arr {
		vInt := int(v)

		if vInt < min {
			min = vInt
		}

		if vInt > max {
			max = vInt
		}

		total += vInt
	}

	fmt.Println(total-max, total-min)
}
