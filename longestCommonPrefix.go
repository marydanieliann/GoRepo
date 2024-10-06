package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for _, str := range strs[1:] {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}

/*func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Common prefix:", longestCommonPrefix(strs))
}*/
