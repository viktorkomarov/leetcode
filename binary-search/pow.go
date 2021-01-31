package binary_search

func myPow(x float64, n int) float64 {
	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}

	return result
}
