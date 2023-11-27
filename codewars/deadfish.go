package codewars

func Parse(data string) []int {
	sum := 0
	var result []int
	for _, v := range data {
		switch {
		case v == 'i':
			sum++
		case v == 'd':
			sum--
		case v == 's':
			sum = sum * sum
		case v == 'o':
			result = append(result, sum)
		}
	}
	return result
}
