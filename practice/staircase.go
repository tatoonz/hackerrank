package practice

import (
	"fmt"
	"strings"
)

// Staircase prints a staircase of size n
// URL: https://www.hackerrank.com/challenges/staircase/problem
func Staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		fmt.Print(strings.Repeat(" ", int(n-i)), strings.Repeat("#", int(i)))
		fmt.Println()
	}
}
