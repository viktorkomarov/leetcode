package march

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	changes := make([]int, amount+1)
	for i := range changes {
		changes[i] = amount + 1
	}

	changes[0] = 0
	for i := 1; i <= amount; i++ {
		for j := range coins {
			if coins[j] <= i {
				changes[i] = min(changes[i], changes[i-coins[j]]+1)
			}
		}
	}

	if changes[amount] > amount {
		return -1
	}

	return changes[amount]
}
