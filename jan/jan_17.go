package leetcode

func lexicGenerate(start string, min rune, alphabet []rune, limit int) int {
	if limit == 0 {
		return 1
	}

	result := 0
	for _, in := range alphabet {
		if in <= min {
			result += lexicGenerate(start+string(min), in, alphabet, limit-1)
		}
	}

	return result
}

func countVowelStrings(n int) int {
	return lexicGenerate("", 'a', []rune{'a', 'e', 'i', 'o', 'u'}, n)
}
