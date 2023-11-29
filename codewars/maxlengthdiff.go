package codewars

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	result := 0
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	for _, item := range a1 {
		for _, a := range a2 {
			result = int(math.Max(float64(result), math.Abs(float64(len(item)-len(a)))))
		}
	}
	return result
}
