package practice_test

import "github.com/tatoonz/hackerrank/practice"

func ExampleMiniMaxSum1() {
	practice.MiniMaxSum([]int32{1, 3, 5, 7, 9})
	// output:
	// 16 24
}

func ExampleMiniMaxSum2() {
	practice.MiniMaxSum([]int32{7, 69, 2, 221, 8974})
	// output:
	// 299 9271
}
