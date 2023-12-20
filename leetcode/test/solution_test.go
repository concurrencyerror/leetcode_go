package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	nums = append(nums[:0], nums[1:]...)
	fmt.Println(nums)
}
