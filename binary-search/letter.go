package binary_search

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters)

	for left < right {
		mid := left + (right-left)/2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return letters[left%len(letters)]
}
