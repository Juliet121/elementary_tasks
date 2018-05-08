package main

import (
	"fmt"

	"strconv"
	"regexp"
)
type context struct {
	min int
	max int
}

var c context

func input1() (string, string) {
	var min, max string
	fmt.Print("Please enter the minimum of the interval to look for Fibonacci numbers:    ")
	fmt.Scanf("%s\n", &min)
	fmt.Print("Please enter the maximum of the interval to look for Fibonacci numbers:    ")
	fmt.Scanf("%s\n", &max)
	return min, max
}

func inputcheck2(min, max string) ( context){
	min1, _ := strconv.Atoi(min)
	max1, _ := strconv.Atoi(max)
	r := regexp.MustCompile("^[0-9]+$")
	if (r.MatchString("min")) != true || (r.MatchString("max")) != true ||
		(r.MatchString("min")) != true && (r.MatchString("max")) != true ||
		min1 < 0 && max1 < 0{
		fmt.Println("Error. You entered a negative number, letter or a number with a separator." +
			"For negative minimum, the program will print the Fibonacci numbers starting from 1(as a default)")
	}
			if max1 < min1 {
				fmt.Println("Min value exceeds the max. Please restart the app by entering the correct data")
			} else {
				c = context{min1, max1}
	}
	return c
}

func main() {
	c:=inputcheck2(input1())
	fmt.Println("Diapason", c, "Fibonacci numbers within this diapason:  ", c.Fibonacci())

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
