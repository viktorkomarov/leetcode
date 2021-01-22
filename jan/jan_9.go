package leetcode

// TimeLimit
/*func countDiff(one, two string) int {
	diff := 0
	for i, r := range one {
		if rune(two[i]) != r {
			diff++
		}
	}

	return diff
}

func helper(beginWord, endWord string, stopList map[string]bool, count int) int {
	if beginWord == endWord {
		return count
	}

	min := 0
	for word, stop := range stopList {
		if !stop && countDiff(beginWord, word) == 1 {
			stopList[word] = true
			currentCase := helper(word, endWord, stopList, count+1)
			if currentCase != 0 {
				if min == 0 || currentCase < min {
					min = currentCase
				}
			}
			stopList[word] = false
		}

	}

	return min
}*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool)
	for _, word := range wordList {
		words[word] = true
	}

	if _, ok := words[endWord]; !ok {
		return 0
	}

	queue := make([]string, 0, len(wordList))
	queue = append(queue, beginWord)

	level := 1
	for len(queue) != 0 {
		for i := len(queue) - 1; i >= 0; i-- {
			currentWord := []rune(queue[i])

			for j, orig := range currentWord {
				for c := 'a'; c <= 'z'; c++ {
					if c == orig {
						continue
					}

					currentWord[j] = c
					attempt := string(currentWord)
					if attempt == endWord {
						return level + 1
					}

					if _, ok := words[attempt]; ok {
						queue = append(queue, attempt)
						delete(words, attempt)
					}
				}
				currentWord[j] = orig
			}

			queue = append(queue[:i], queue[i+1:]...)
		}

		level++
	}

	return 0
}
