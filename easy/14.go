package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	lenCache := make([]int, len(strs))
	maxLen := 0
	for index, str := range strs {
		if str == "" {
			return ""
		}
		if str[0] != strs[0][0] {
			return ""
		}
		lenCache[index] = len(str)
		if len(str) >= maxLen {
			maxLen = len(str)
		}
	}

	for index := 1; index < maxLen; index++ {
		for i, str := range strs {
			if index >= lenCache[i] || str[index] != strs[0][index] {
				return strs[0][:index]
			}
		}
	}

	return strs[0]
}
