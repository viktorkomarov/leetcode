package leetcode

func concatenatedBinary(n int) int {
	var ans int64

	for i := 1; i <= n; i++ {
		size := digits(i)
		ans = ((ans << size) | int64(i)) % (1e9 + 7)
	}

	return int(ans)
}

func digits(i int) int {
	shift := 31
	for ; (1<<shift)&i == 0; shift-- {
	}

	return shift + 1
}
