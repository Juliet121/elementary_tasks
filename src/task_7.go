package main

import (
	"fmt"

)
type context struct {
	min int
	max int
}

func main() {
	var c context
	var max, min int

	/*fmt.Print("Please enter the minimum of the interval to look for Fibonacci numbers:    ")
	fmt.Scanf("d\n", min)
	fmt.Print("Please enter the maximum of the interval to look for Fibonacci numbers:    ")
	fmt.Scanf("d\n", max)*/
 min = 85
 max = 1800
	if c.min < 0 && c.max < 0 {
		fmt.Println("Please enter the positive integer number")
	} else {
		c = context{min, max}
		fmt.Println(c)
	}
	fmt.Println("Fibonacci numbers on the diapason:  ", c.Fibonacci())
}

func (c *context) Fibonacci() (f[]int) {
	var a, b int
	a = 0
	b = 1
	for i:=c.min; i < c.max; i++ {
		a, b = b, a+b
		if b < c.max && b > c.min {
		f = append(f, b)
		}
	}
		return f

}
