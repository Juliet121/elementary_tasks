package main

import (
	"fmt"
)

func main() {
	var height, width int
	fmt.Print("Could you enter the height of chess desk please: ")
	fmt.Scanf("%d", &height)
	fmt.Print("Could you enter the width of chess desk please: ")
	fmt.Scanf("%d", &width)

	if height > 0 && width > 0 {
		symbol:="*"
		j:=chessdesk(height, width, symbol)
		fmt.Println(j)

	} else {
		fmt.Println("User entered incorrect data")
	}
}

func chessdesk(height, width int, symbol string) (str string){
	str=" "
	for i:=0; i < width; i++{
		//we are adding additional space between rows:
		if i%2==0 {
			str+=" "
		}

		for j:=0; j < height; j++ {
			str+=" " + symbol + " "
		}
		str+="\n"
	}
	return str
}



