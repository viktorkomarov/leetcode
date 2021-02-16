package feb

import "unicode"

func letterCasePermutation(S string) []string {
	result := []string{""}

	for _, r := range S {
		if unicode.IsDigit(r) {
			for j := range result {
				result[j] += string(r)
			}
		} else {
			small := string(unicode.ToLower(r))
			big := string(unicode.ToUpper(r))

			update := make([]string, 0)
			for j := range result {
				update = append(update, result[j])
				result[j] += small
			}

			for _, str := range update {
				result = append(result, str+big)
			}
		}
	}

	return result
}
