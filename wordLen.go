package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}

/*func main() {
	fmt.Println(lengthOfLastWord("Hello World"))                 // Output: 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // Output: 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // Output: 6
}*/
