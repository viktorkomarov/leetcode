package leetcode

func getSmallestString(n int, k int) string {
	alphabet := make([]rune, 0)
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, i)
	}
	//'abcdefghijklmnopqrstuvwxyz'
	result := make([]rune, n)
	for i := n - 1; i >= 0; i-- {
		symbol := (k - i)
		if symbol >= len(alphabet) {
			symbol = len(alphabet)
		}
		k -= symbol
		result[i] = alphabet[symbol-1]
	}

	return string(result)
}
