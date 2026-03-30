package main

import (
	"fmt"
	//"strings"

)

func palindrome(s string) string {
	rev := []rune(s)
	var reverse []rune
	for i := len(rev) - 1; i >= 0; i-- {
		reverse = append(reverse, rev[i])
	}
	if string(reverse) == string(rev) {

		return "is a palindrome"
	}
	return "is not a palindrome"
}


	


func main() {
	fmt.Println(palindrome("neveroddoreven"))
}
