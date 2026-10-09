func longestCommonPrefix(strs []string) string {
	res := ""
	minSize := 9999999999

    for i := 0; i < minSize;i++{
		for j := 0; j < len(strs); j++ {
			if strs[j] == "" {
				return ""
			}
			if len(strs[j]) < i || strs[j][i] != strs[0][i] {
				return res
			}

			minSize = min(minSize, len(strs[j]))
		}

		res = res + string(strs[0][i])
	}

	return res
}
