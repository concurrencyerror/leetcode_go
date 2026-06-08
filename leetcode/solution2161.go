package leetcode

import "slices"

func pivotArray(nums []int, pivot int) []int {
	big := make([]int, 0)
	small := make([]int, 0)
	middle := make([]int, 0)
	for _, num := range nums {
		if num < pivot {
			small = append(small, num)
		} else if num > pivot {
			big = append(big, num)
		} else if num == pivot {
			middle = append(middle, num)
		}
	}
	return slices.Concat(small, middle, big)
}
