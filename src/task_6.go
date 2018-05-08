package main

import ("fmt"
	"strconv"
	"regexp"
)

func input () (string, string){
	var lenght, minSq string
	fmt.Println("Please enter the lenght of the string of natural numbers")
	fmt.Scanf("%s\n", &lenght)
	fmt.Println("Please enter the number m - minimal square")
	fmt.Scanf("%s\n", &minSq)
	return lenght, minSq
}

func checkthedata (lenght string, minSq string) (lenght1 int, minSq1 int) {
	lenght1, _ = strconv.Atoi(lenght)
	minSq1, _ = strconv.Atoi(minSq)
	r := regexp.MustCompile("^[0-9]+$")

	if (r.MatchString("minSq")) != true || (r.MatchString("lenght")) != true ||
		(r.MatchString("lenght")) != true && (r.MatchString("minSq")) != true {
		fmt.Println("Error. You entered a letter or a number with a separator.Program exited")
	}
	if lenght1 < 0 || minSq1 < 0{
		fmt.Println("You entered the negantive number. Please restart the program.")
	} else{
		getSquares(lenght1, minSq1)
	}
	return
}


func main() {
	seq := Printwithcommas(getSquares(checkthedata(input())))
	fmt.Println("Numerical sequence separated by commas:  ", seq)
}

func getSquares(length1, minSq1 int) []int {
	var sq int
	j := 0
	arr := []int{}
	for  {
		sq = j*j
		if sq > minSq1 {
			if len(arr) == length1 {
				break;
			}
			arr = append(arr, sq)
		}
		j++
	}
	return arr
}

func Printwithcommas (arr []int) string {
		if len(arr) == 0 {
			return ""
		}
		estimate := len(arr) * 4
		b := make([]byte, 0, estimate)
			for _, n := range arr {
				b = strconv.AppendInt(b, int64(n), 10)
				b = append(b, ',')
			}
				b = b[:len(b)-1]
	return string(b)
}







