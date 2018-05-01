package main
import (
	"fmt"
	_"strconv"
	"strings"
)

func main() {
	var str string

	fmt.Println("Please enter a sequence consisting of not less than 2 digits/signs " +
		"Be aware that the number should not contain any separators     ")
	fmt.Scanf("%s\n", &str)

	if len(str) < 2 {
		fmt.Println("You entered a sequence consisting of less than 2 digits. Program exited")
	} else {
		if str == Reverse(str) {
			fmt.Println(str, "The sequence is a palindrome")
		}
	}
	palindromes := []string{}

	for i := 0; i < len(str); i++ {
		ch := string(str[0:i])
		if strings.Index(str[i+1:(len(str))], Reverse(ch)) != -1 {
			palindromes = append(palindromes, ch)
		} else {
			fmt.Println("0.No more palindromes were found")
			break
			}
		fmt.Println(palindromes)
	}
}




//Reverse is used to reverse the string, that afterwards will be compared with the initial string
func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}