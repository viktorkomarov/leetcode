package leetcode

func mostCompetitive(n []int, k int) []int {
	s := make([]int, 0)
	for i, _ := range n {
		for len(s) > 0 && s[len(s)-1] > n[i] && (len(s)+len(n)-i) > k {
			s = s[:len(s)-1]
		}
		s = append(s, n[i])
	}
	return s[:k]
}
