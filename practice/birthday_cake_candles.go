package practice

// BirthdayCakeCandles returns tallest candles that can successfully blow out
// URL: https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func BirthdayCakeCandles(ar []int32) int32 {
	countMap := map[int32]int{}
	tallest := int32(0)

	for _, v := range ar {
		countMap[v]++

		if v > tallest {
			tallest = v
		}
	}

	return int32(countMap[tallest])
}
