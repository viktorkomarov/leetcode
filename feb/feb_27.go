package feb

func divide(dividend int, divisor int) int {
	if -1<<31 == dividend && divisor == -1 {
		return 1<<31 - 1
	}
	var quotient int
	sign := 1
	if (dividend > 0) != (divisor > 0) {
		sign = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}
	if sign == 1 {
		return quotient
	}
	return -quotient
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
