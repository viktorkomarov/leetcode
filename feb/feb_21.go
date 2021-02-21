package feb

func brokenCalc(X int, Y int) int {
	result := 0

	for Y > X {
		result++
		if Y%2 == 1 {
			Y++
		} else {
			Y /= 2
		}
	}

	return result + X - Y
}
