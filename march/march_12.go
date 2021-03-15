package march

func hasAllCodes(s string, k int) bool {
	dir := make(map[int]bool)

	for i := 0; i <= len(s)-k; i++ {
		num,_ := strconv.ParseInt(s[i:i+k],2,64)
		dir[int(num)] = true
	}

	for i := 0; i < (k << 1); i++ {
		if !dir[i] {
			return false
		}
	}

	return true
}

