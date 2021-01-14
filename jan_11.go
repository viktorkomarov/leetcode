package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for _, num := range nums2 {
		nums1[m] = num
		for i := m - 1; i >= 0 && num < nums1[i]; i-- {
			nums1[i], nums1[i+1] = num, nums1[i]
		}
		m++
	}
}
