package main

import (
	"reflect"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	if x < 0 {
		rev += "-"
	}
	if reflect.DeepEqual(str, rev) {
		return true
	} else {
		return false
	}
}

/*func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-123))
}
*/
