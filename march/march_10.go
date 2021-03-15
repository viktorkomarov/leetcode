package march

import "strings"

func romans(order int, val int) string {
	nums := []rune{
		'I', 'V', 'X', 'L',
		'C', 'D', 'M',
	}

	order *= 2

	if val == 0 {
		return ""
	}

	if val <= 3 {
		temp := ""
		for i := 0; i < val; i++ {
			temp += string(nums[order])
		}

		return temp
	}

	if val == 4 {
		return string(nums[order]) + string(nums[order+1])
	}

	if val == 5 {
		return string(nums[order+1])
	}

	if val <= 9 {
		temp := string(nums[order+1])
		for i := 0; i < val-5; i++ {
			temp += string(nums[order])
		}

		return temp
	}

	return string(nums[order]) + string(nums[order+2])
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func intToRoman(num int) string {
	acc := make([]string, 0)
	order := 0

	for ; num > 0; num /= 10 {
		n := num % 10
		acc = append(acc, romans(order, n))

		order++
	}

	return strings.Join(reverse(acc), "")
}
