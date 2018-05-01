package main

import "fmt"

func main() {
	var lenght, minSq int
	fmt.Println("Please enter the lenght of the string of natural numbers")
	fmt.Scan(&lenght)
	fmt.Println("Please enter the number m - minimal square")
	fmt.Scan(&minSq)
	if lenght < 0 || minSq < 0 {
		fmt.Println("Error. You entered negative numbers")
	}

	arr := getSquares(lenght, minSq)
	fmt.Println(arr)
}
func getSquares(length, minSq int) []int {
	var sq int
	j := 0
	arr := []int{}
	for  {
		sq = j*j
		if sq > minSq {
			if len(arr) == length {
				break;
			}
			arr = append(arr, sq)
		}
		j++
	}
	return arr
}




