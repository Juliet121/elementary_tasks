package main

import "fmt"

type Envelope struct {
	Height float64
	Width float64
}

func compareEnvelopes(e1, e2 Envelope) int {
	res := 0
	if e1.Width >= e2.Width && e1.Height >= e2.Height {
		res = 1
		fmt.Println("It's possible to put envelope 2 into envelope 1")
	}
	if e1.Width < e2.Width && e1.Height < e2.Height {
		res = 2
		fmt.Println("It's possible to put envelope 1 into envelope 2")
	}
	if e1.Width == e2.Height && e1.Height == e2.Width {
		res = 1
		fmt.Println("If you swap either of two envelopes, they can be put one into another")
	}

	//проверка возможности размещения конверта под углом
	if e1.Width*e1.Height > 2*e2.Width*e2.Height && (e2.Width*e2.Width+e2.Height*e2.Height-e1.Height*e1.Height)*(e2.Width*e2.Width+e2.Height*e2.Height-e1.Width*e1.Width) <=
		(e1.Width*e1.Width*e1.Height*e1.Height - 4*e1.Width*e1.Height*e2.Height*e2.Width + 4*e2.Width*e2.Width*e2.Height*e2.Height) {
		res = 1
		fmt.Println("The envelope can be placed diagonally into the envelope #", res)
	}
	return res
}
func main() {
	var e1, e2 Envelope

	fmt.Print("Please enter the height of the first envelope: ")
	fmt.Scanf("%f\n", &e1.Height)

	fmt.Print("Please enter the width of the first envelope:  ")
	fmt.Scanf("%f\n", &e1.Width)

	fmt.Print("Please enter the height of the second envelope:  ")
	fmt.Scanf("%f\n", &e2.Height)

	fmt.Print("Please enter the width of the second envelope:  ")
	fmt.Scanf("%f\n", &e2.Width)

	if e1.Height > 0 && e1.Width > 0 && e2.Height > 0 && e2.Width > 0 {
		s := compareEnvelopes(e1, e2)
		fmt.Println(s)
	} else {
		fmt.Println("User entered incorrect data. Please check if you use positive number")
	}
}

func printValue(Envelope) (e1, e2 Envelope){
	e1 = Envelope{e1.Height, e1.Width}
	e2 = Envelope{e2.Height, e2.Width}
	fmt.Println(e1, e2)
	return e1, e2
}