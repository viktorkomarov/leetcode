package feb

func isArithmetic(num []int) (int, bool) {
	step := num[1] - num[0]
	if num[2]-num[1] != step {
		return 0, false
	}

	return step, true
}

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	count := 0
	for i := 0; i < len(A)-2; i++ {
		step, ok := isArithmetic(A[i : i+3])
		if !ok {
			continue
		}

		count++
		for j := i + 3; j < len(A); j++ {
			if A[j]-A[j-1] == step {
				count++
			} else {
				break
			}
		}
	}

	return count
}
