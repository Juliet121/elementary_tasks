	package main

	import (
		"fmt"
		"strconv"
	)

func inputtocheck() (height, width string){

	fmt.Printf("%s\n", "Could you enter the height of chess desk please: ")
	fmt.Scanf("%s\n", &height)
	fmt.Printf("%s\n","Could you enter the width of chess desk please: ")
	fmt.Scanf("%s\n", &width)
	height1,_ := strconv.Atoi(height)
	width1,_ := strconv.Atoi(width)
	if height1 > 0 && width1 > 0 {
		symbol:="*"
		j:=chessdesk(height1, width1, symbol)
		fmt.Printf("%s\n",j)
	} else {
		fmt.Printf("%s\n","User entered incorrect data. Probably negative digit or a digit" +
			" with separators. Program is working appropriately only with integers. Please re-start it")
			}
	return

}

func main() {
	inputtocheck()
}

func chessdesk(height1, width1 int, symbol string) (str string){
		str=" "
		for i := 0; i < width1; i++{
		//we are adding additional space between rows:
			if i%2==0 {
				str+=" "
			}
			for j := 0; j < height1; j++ {
				str+=" " + symbol + " "
			}
			str+="\n"
		}
		return str
}



