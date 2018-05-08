package main

import (
	"fmt"
	"math"
)

	type Envelope struct {
		height float64
		width float64
	}

func validateinput() (e1, e2 Envelope) {

	fmt.Print("Please enter the height of the first envelope: ")
	fmt.Scanf("%f\n", &e1.height)

	fmt.Print("Please enter the width of the first envelope:  ")
	fmt.Scanf("%f\n", &e1.width)

	fmt.Print("Please enter the height of the second envelope:  ")
	fmt.Scanf("%f\n", &e2.height)

	fmt.Print("Please enter the width of the second envelope:  ")
	fmt.Scanf("%f\n", &e2.width)
	return e1, e2
}

func check(e1, e2 Envelope) (){
	if e1.height > 0 && e1.width > 0 && e2.height > 0 && e2.width > 0 {
		s := compareEnvelopes(e1, e2)
		fmt.Println(s)
	} else {
			fmt.Println("User entered incorrect data. Please check if you use positive number")
			}
}


func compareEnvelopes(e1, e2 Envelope) int {
	res := 0
		if e1.width >= e2.width && e1.height >= e2.height {
			res = 1
			fmt.Println( "It's possible to put envelope 2 into envelope 1")
		}
			if e1.width < e2.width && e1.height < e2.height {
				res = 2
				fmt.Println("It's possible to put envelope 1 into envelope 2")
			}
				if e2.width > e1.width &&
					e1.height >= (2*e2.width*e2.height*e1.width+((math.Pow(e2.width, 2)-(math.Pow(e2.height, 2))*math.Sqrt(
					(math.Pow(e2.width, 2)+math.Pow(e2.height, 2)-(math.Pow(e1.width, 2)))))))/(math.Pow(e2.width, 2)+
					math.Pow(e2.height,2)) {
					res = 2
					fmt.Println("If turned, the envelope 1 can be placed into the envelope #", res)
				}
					if	e1.width > e2.width &&
						e2.height >=(2*e1.width*e1.height*e2.width+((math.Pow(e1.width, 2)-(math.Pow(e1.height, 2))*math.Sqrt(
						(math.Pow(e1.width, 2)+math.Pow(e1.height, 2)-(math.Pow(e2.width, 2)))))))/(math.Pow(e1.width, 2)+
							math.Pow(e1.height, 2)){
						res = 1
						fmt.Println("If turned, the envelope 2 can be placed into the envelope #", res)
					}
					return res
}

func main() {
		check(validateinput())
}


	

