package leetcode_go

func averageValue(nums []int) int {
	result := 0
	i := 0
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			result += num
			i++
		}
	}
	if result == 0 {
		return 0
	}
	return result / i
}
