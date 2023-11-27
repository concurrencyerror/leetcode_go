package codewars

import "math"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	target := math.Sqrt(float64(n))
	for i := 2; i <= int(target); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
