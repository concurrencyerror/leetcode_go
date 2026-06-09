package leetcode

import "math"

func maxTotalValue(nums []int, k int) int64 {
	numMin, numMax := math.MaxInt, math.MinInt
	for _, num := range nums {
		numMin = min(numMin, num)
		numMax = max(numMax, num)
	}
	return int64((numMax - numMin) * k)
}
