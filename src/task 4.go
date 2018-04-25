package main

import (
	"fmt"
	/*s "strings"*/
	"strconv"
)


func main() {var n, l int
	var n_str string
	fmt.Println("Please enter a positive number consisting of not less than 10 digits. Be aware that the number should not contain any separators(like commas, points)     ")
	fmt.Scanf("%d\n", &n)
	n_str = strconv.Itoa(n)
	l = len(n_str)

	if l < 10 {
		fmt.Println("Please try again and enter at least 10-digit number")
		fmt.Scanf("%d\n", &n)
	} else {

		s := make([]string, 0)
		for i := 0; i < l; i++ {
			s = append(s, s[i])
		}
		fmt.Println(s)
		Palindrome(s,l)
	}
}

func Palindrome([]string, int) (string, int) {
	for i := 0; i < l/2; i++ {
		if s[i] == s[(l-1)-i] {
			j := "Congratulations! You found the palindrome"
		} else {
			j := "Sorry, this number is not a palindrome "
		}
	}
	return
}


