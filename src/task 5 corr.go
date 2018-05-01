package main

import (
	"fmt"
	"strconv"
	_"strings"
	"strings"
)

type Ticket struct{
	Min, Max int
	Value string
}

type Winner struct{
	Method string
	Count1, Сount2 int
}

func (t *Ticket) Simple() (isHappy bool){
	var sumVal1, sumVal2 int

	count := 6
	isHappy = false
	sArr := strings.Split(t.Value, "")

	if len(t.Value) == count {
		sumVal1 = Sum(sArr[0:count/2])
		sumVal2 = Sum(sArr[count/2: count])

		if (sumVal1 == sumVal2) {
			isHappy = true
		}
	}

	return
}

func (t *Ticket) Difficult() (isHappy bool) {
	odd := []string{}
	even := []string{}
	sArr := strings.Split(t.Value, "")
	for i := 0; i < len(sArr); i++ {
		n, _ := strconv.Atoi(sArr[i])
		if n % 2 == 0 {
			even =append(even, sArr[i])
		} else {
			odd = append(odd, sArr[i])
		}
	}

	sumVal1 := Sum(odd)
	sumVal2 := Sum(even)


	if (sumVal1 == sumVal2) {
		isHappy = true
	}

	return
}

func Sum(str []string) (s int) {
	s = 0;
	for i := 0; i < len(str); i++ {
		n, _ := strconv.Atoi(string(str[i]))
		s += n
	}

	return
}

func Compare(t *Ticket) (w Winner){
	count1 :=0
	count2 :=0
	for i:=t.Min; i < t.Max; i++{
		t.Value = strconv.Itoa(i)
		happy1:=t.Simple()
		if happy1{
			count1++
		}
		happy2:=t.Difficult()
		if happy2{
			count2++
		}
	}
	if count1 > count2 {
		w.Method="Simple"
	} else {
		w.Method="Difficult"
	}
	w.Count1 = count1
	w.Сount2 = count2
	fmt.Println("count1", count1, "count2", count2)
	return
}


func main() {

	t := Ticket{}
	fmt.Print("Please enter the minimum of the interval:    ")
	fmt.Scanf("%d\n", &t.Min)
	fmt.Print("Please enter the maximum of the interval:    ")
	fmt.Scanf("%d\n", &t.Max)

	t.Value = strconv.Itoa(t.Min)

	happy := t.Simple()
	fmt.Print(happy)
	happy = t.Difficult()
	fmt.Println(happy)

	fmt.Println(Compare(&t))


}