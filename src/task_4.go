package main
import (
	"fmt"
	_"strconv"
	"strings"
	"unicode/utf8"
)

func Inputvalid() (string){
	var str string
	fmt.Printf("%s\n", "Please enter a sequence consisting of not less than 2 digits/signs " +
		"Be aware that the number should not contain any separators     ")
	fmt.Scanf("%s\n", &str)
		if utf8.RuneCountInString(str) < 2 {
			fmt.Printf("%s\n", "You entered a sequence consisting of less than 2 digits. Program exited")
				} else {
					if str == Reverse(str) {
					fmt.Println( str, "The sequence is a palindrome")
					}
				}
	return str
}

func PalindromesSearch(str string) ([]string){
	palindromes := []string{}
		for i := 0; i < utf8.RuneCountInString(str); i++ {
			ch := string(str[0:i])
				if strings.Index(str[i+1:(utf8.RuneCountInString(str))], Reverse(ch)) != -1 {
					palindromes = append(palindromes, ch)
					} else {
					fmt.Printf("%s\n","0.No more palindromes were found")
					break
					}
					fmt.Printf("%s\n", palindromes)
		}
		return palindromes
}

func main() {
	PalindromesSearch(Inputvalid())
}

//Reverse is used to reverse the string, that afterwards will be compared with the initial string
func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}