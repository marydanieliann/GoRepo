package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

/*func main() {
	testCases := []struct {
		haystack string
		needle   string
	}{
		{"sadbutsad", "sad"},
		{"leetcode", "leeto"},
		{"hello", "ll"},
		{"", ""},
		{"abc", "abc"},
	}

	for _, testCase := range testCases {
		result := strStr(testCase.haystack, testCase.needle)
		fmt.Printf("Input: haystack = %q, needle = %q, Output: %d\n", testCase.haystack, testCase.needle, result)
	}
}
*/
