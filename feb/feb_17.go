package feb

func min(l, r int) int {
	if l < r {
		return l
	}

	return r
}

func maxArea(height []int) int {
	max := -1
	l := 0
	r := len(height) - 1

	for l < r {
		sq := min(height[l], height[r])
		if sq*(r-l) > max {
			max = sq * (r - l)
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}
