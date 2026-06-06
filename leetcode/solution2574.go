package leetcode

import "math"

func leftRightDifference(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	result := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			left[i] = 0
		} else {
			left[i] = left[i-1] + nums[i-1]
		}
	}
	for t := len(nums) - 1; t >= 0; t-- {
		if t == len(nums)-1 {
			right[t] = 0
		} else {
			right[t] = nums[t+1] + right[t+1]
		}
	}
	for i := range result {
		result[i] = int(math.Abs(float64(left[i] - right[i])))
	}
	return result
}
