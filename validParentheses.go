package main

func isValid(s string) bool {
	//data type to represent Unicode code point
	//alias for int32
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if openBracket, exists := bracketMap[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

/*
func main() {
	sample := []string{
		"()", "()[]{}", "(]", "([])", "{[()]}", "{[(])}", "{{[[(())]]}}", "(()",
	}

	for _, s := range sample {
		result := isValid(s)
		fmt.Printf("Input: %s, Output: %v\n", s, result)
	}
}
*/
