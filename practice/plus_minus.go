package practice

import "fmt"

// PlusMinus calculate the fractions of its elements that are positive, negative, and are zeros and print the decimal value of each fraction on a new line
// URL: https://www.hackerrank.com/challenges/plus-minus/problem
func PlusMinus(arr []int32) {
	countFraction := [3]int{}

	for _, n := range arr {
		if n > 0 {
			countFraction[0]++
		} else if n < 0 {
			countFraction[1]++
		} else {
			countFraction[2]++
		}
	}

	for _, c := range countFraction {
		fmt.Printf("%.6f\n", float64(c)/float64(len(arr)))
	}
}
