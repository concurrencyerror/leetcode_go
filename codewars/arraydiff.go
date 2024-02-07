package codewars

func ArrayDiff(a, b []int) []int {
	if len(a) == 0 {
		return a
	}
	result := make([]int, 0)
	for _, i := range a {
		if !contains(b, i) {
			result = append(result, i)
		}
	}
	return result
}

func contains(b []int, value int) bool {
	for _, i := range b {
		if i == value {
			return true
		}
	}
	return false
}
