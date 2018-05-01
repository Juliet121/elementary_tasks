package main

import "fmt"

func main() {





	fmt.Println(Average(19, 24, 98, 76, 85, 10))
}

func Average(nums...int) int {
	fmt.Print(nums, " ")
	r := []int{}
	summa:=0
	for i := 0; i <len(nums); i++ {
		summa=summa+nums[i]
		r =append(r, nums[i])
	}
	d:=summa/len(r)
	return d
}










/*fmt.Println("len:  ", len("Julia Saryan"))
fmt.Println("Char: ", "Yuliya"[5])
fmt.Print("ToUpper:  ", s.ToUpper("Yuliya"))
fmt.Println("Contains ", s.Contains("Yuliya ", "y"))
fmt.Println("Repeat  ", s.Repeat("Yuliya ", 5))
fmt.Println("Split:     ", s.Split("a-b-c-d-e", " -"))*/

	//type Envelope struct {
	//		Height float64
	//		Width float64
	//	}

	//	func compareEnvelopes(e1, e2 Envelope) int {
	//		res := 0
	//		if e1.Width >= e2.Width && e1.Height >= e2.Height {
	//			res = 1
	//			fmt.Println("It's possible to put envelope 2 into envelope 1")
	//		}
	//		if e1.Width < e2.Width && e1.Height < e2.Height {
	//			res = 2
	//			fmt.Println("It's possible to put envelope 1 into envelope 2")
	//		}
	//		if e1.Width == e2.Height && e1.Height == e2.Width {
	//			res = 1
	//			fmt.Println("If you swap either of two envelopes, they can be put one into another")
	//		}

	//проверка возможности размещения конверта под углом
	//		if e1.Width*e1.Height > 2*e2.Width*e2.Height && (e2.Width*e2.Width+e2.Height*e2.Height-e1.Height*e1.Height)*(e2.Width*e2.Width+e2.Height*e2.Height-e1.Width*e1.Width) <=
	//			(e1.Width*e1.Width*e1.Height*e1.Height - 4*e1.Width*e1.Height*e2.Height*e2.Width + 4*e2.Width*e2.Width*e2.Height*e2.Height) {
	//			res = 1
	//			fmt.Println("The envelope can be placed diagonally into the envelope #", res)
	//		}
	//		return res
	//	}
	//	func main() {
	//		var e1, e2 Envelope
	//
	//		fmt.Print("Please enter the height of the first envelope: ")
	//		fmt.Scanf("%f\n", &e1.Height)

	//		fmt.Print("Please enter the width of the first envelope:  ")
	//		fmt.Scanf("%f\n", &e1.Width)

	//		fmt.Print("Please enter the height of the second envelope:  ")
	//		fmt.Scanf("%f\n", &e2.Height)

	//		fmt.Print("Please enter the width of the second envelope:  ")
	//		fmt.Scanf("%f\n", &e2.Width)

	//		if e1.Height > 0 && e1.Width > 0 && e2.Height > 0 && e2.Width > 0 {
	//			s := compareEnvelopes(e1, e2)
	//			fmt.Println(s)
	//		} else {
	//			fmt.Println("User entered incorrect data. Please check if you use positive number")
	////		}
	//	}

	//		func printValue(Envelope) (e1, e2 Envelope){
	//			e1 = Envelope{e1.Height, e1.Width}
	//			e2 = Envelope{e2.Height, e2.Width}
	//			fmt.Println(e1, e2)
	//			return e1, e2
	//		}

	//----------------------------

	/*type Triangle struct {
	Vertices string
	Side_A float64
	Side_B float64
	Side_C float64
}

func (c *Triangle) halfperimeter() (p float64) {
	p = (c.Side_C+c.Side_B+c.Side_A)/2
	return p
}


func (c *Triangle) area() (a float64) {
	a = math.Sqrt(c.halfperimeter()*(c.halfperimeter()-c.Side_A)*(c.halfperimeter()-c.Side_B)*(c.halfperimeter()-c.Side_C))
	return
}

func ComparisonofTriangles(triangles []Triangle) (t []string) {
	//bubble sorting
	for i:=1; i< len(triangles); i++ {
		for j:=0; j < len(triangles)-i; j++ {
			if  triangles[j].area() < triangles[j+1].area() {
				triangles[j], triangles[j+1] = triangles[j+1], triangles[j]
			}
		}
	}

	for i:=0; i< len(triangles); i++ {
		t = append(t, triangles[i].Vertices)
	}
	return
}


func main() {

	triangles := []Triangle{}
	for i:=0; i < 3; i++ {
		t1 := Triangle{}
		fmt.Println("Please enter the symbolic name of vertices of the " , i+1 , " triangle:    ")
		fmt.Scanf("%s", &t1.Vertices)
		fmt.Println("Please enter side A of the " , i+1 , " triangle:     ")
		fmt.Scan(&t1.Side_A)
		fmt.Println("Please enter side B of the " , i+1 , "triangle:     ")
		fmt.Scan(&t1.Side_B)
		fmt.Println("Please enter side C of the " , i+1 , " triangle:     ")
		fmt.Scan (&t1.Side_C)
		if t1.Side_A > 0 && t1.Side_B > 0 && t1.Side_C > 0 {
			triangles = append(triangles, t1)
		} else {
			fmt.Println("User entered incorrect input data (negative digit)")
			break
		}
	}
	t:=ComparisonofTriangles(triangles)
	fmt.Println("Title of triangles: ", t)
}

*/

	///1111
	//	var height, width int
	//	fmt.Print("Could you enter the height of chessdesk please: ")
	//	fmt.Scanf("%d", &height)
	//	fmt.Print("Could you  enter the width of chessdesk please: ")
	//	fmt.Scanf("%d", &width)

	//	if height > 0 && width > 0 {
	//		symbol:="*"
	//		j:=chessdesk(height, width, symbol)
	//		fmt.Println(j)

	//	} else {
	//		fmt.Println("User entered incorrect data")
	//}
	//}

	//func chessdesk(height, width int, symbol string) (str string){
	//	str=" "
	//	for i:=0; i < width; i++{
	//		//we are adding additional space between rows:
	//		if i%2==0 {
	//			str+=" "
	//		}
	//		for j:=0; j < height; j++ {
	//			str+=" " + symbol + " "
	//		}
	//		str+="\n"
	//	}
	//	return str
	//}

	//1

	//import (
	//"math"
	//)

	//type Triangle struct {
	//	vertices string
	//	a float64
	//	b float64
	//	c float64
	//}

	//--------------
	//*Конверты
	//type Envelope struct {
	//	Height float64
	//	Width float64
	//}
	//func (r *Envelope) square() float64 { метод
	//	return r.Width*r.Height
	//}

	//func compareEnvelopes(e1, e2 Envelope) int {
	//	res := 0
	//	if e1.Width > e2.Width && e1.Height > e2.Height {
	//		res = 1
	//	}
	//	if e1.Width < e2.Width && e1.Height < e2.Height {
	//		res = 2
	//	}
	//	return res
	//}

	//func main() {

	//t1 := Triangle{"ABC", 10, 20, 22.36}
	//t2 := Triangle{"DEF", 15, 46.05, 22.01}
	//t3 := Triangle{"KLM", 8, 7, 87.41}
	//var n int
	//var arr[] *Triangle
	//for i := 0; i <= n; i++{
	//	arr = append(arr, t)
	//}
	//t := AreaofTriangles(arr)
	//fmt.Println(t1, t2, t3)
	//}
	//	func AreaofTriangles(arr[]*Triangle) string{
	//str := " "
	//var n int
	//var p,s []float64

	//for i := 0; i <= n; i++{
	//p[i] = (i.a+i.b+i.c)/2
	//s[i] = math.Sqrt(p[i]*(p[i]-a[i])*(p[i]-b[i])*(p[i]-c[i]))
	//fmt.Println(s[i])
	//}
	//for _, r := range (s){
	//str+="s[i]+vertices"
	//str+="\n"
	//}
	//return str
	//}
	//	s:=variadicReverse(2, 3, 5, 5, 3, 6, 11, 11, 87, 46)
	//	fmt.Println(s)
	//}

	//func variadicReverse(nums...int) []int {
	//	fmt.Print(nums, " ")
	//	r := []int{}
	//	for i := len(nums)-1; i >= 0; i-- {
	//	r =append(r, nums[i])
	//	}
	//				return r
	//	}

	//__________________

	//создаем массив из 5 чисел, тип int - 1 способ
	//	var x [5]int16
	//	x[0] = 13
	//	x[1] = 8
	//	x[2] = -9
	//	x[3] = 18
	//	x[4] = 98
	//	fmt.Println("array x", x)

	//инициализируем массив - 2 способ
	//	x1 := []int16{2, 57, 4, 18, 46}
	//	fmt.Println("array x1", x1)

	//создадим срез из массива х1, элементы которого с нулевого по третий
	//	sx1 := x1[0:3]
	//	fmt.Println("slice sx1", sx1)

	//инициализация среза (создание массива не требуется), срез сам создает массив
	//	sx2 := []int{2, 8, -17, 6, 80, -21}
	//	fmt.Println("slice sx2", sx2)

	//append - добавление данных в срез - по умолчанию данніе добавляются в конец массива
	//	sx3 := append(sx2, -9, 80)
	//	fmt.Println("append", sx3)

	//копирование среза в новый - для начала надо создать пустой срез, функция copy его не создает. При этом длина среза
	//должна соответствовать длине копируемого среза - len
	//	sx4 := make([] int, len(sx3))
	//	copy(sx4, sx3)
	//	fmt.Println("copy", sx4)

	//создаем двухмерный срез - вначале инициализируем пустой двухмерный срез, потом с помощью 2 циклов создаем нужные выборки
	//	sd1 := make([][] int, 6)
	//	for i := 0; i < 5; i++ {
	//		innerlength := i + 1
	//		sd1[i] = make([]int, innerlength)
	//		for j := 0; j < innerlength; j++ {
	//			sd1[i][j] = i + j
	//		}

	//	}
	//	fmt.Println(sd1)
	// cоздаем двухмерный массив - выделяя под него память - размер - основное свойство массива.
	//Создали 3 массива по 4 числа

	//	var dd2 [3][4]int
	//	for i := 0; i < 3; i++ {
	//		for j := 0; j < 4; j++ {
	//			dd2[i][j] = i + j
	//		}
	//	}
	//fmt.Println("two-dimensional array", dd2)

	//Gets all but the last element of array.
	//а вернет  мне функция  входной массив кроме  последнего элемента
	//[1,2]
	//	arr := [3]int{1, 2, 3}
	//	func initial(arr []int{}) []int{} {
	//		sl: = arr[0:(len(arr)-1)]
	//		return
	//	}
	//	res := sl
	//	fmt.Println("array lenght", len(arr), "slice", res)

	//	x5 := initial(sx2)
	//	fmt.Println(sx2)
	//	fmt.Println(x5)

	//}
	//func initial(x []int) []int {
	//	return x[0:(len(x)-1)]
	//}

	//	z:=[]int{1,2,1,2}
	//	value:=2
	//	z1:=lastIndexOf(z, value)
	//	fmt.Println(z)
	//	fmt.Println(z1)

	//func lastIndexOf(z []int, value int) int {
	//	ind := -1
	//	for i := 0; i < len(z); i++ {
	//		if z[i] == value{
	//			ind = i
	//		}
	//	}
	//	return ind

	//	r := []int{1, 2, 3, 4, 5}
	//	r1 := last(r)
	//	fmt.Println(r)
	//	fmt.Println(r1)
	//}

	//func last(r[]int) int{
	//	i:=(len(r)-1)
	//	l:=r[i]
	//		return l
	//}

	//	without([1,2,3,4,2,4,2], 4) => [1,2,3,2,2]
	//	r1 := []int{1, 2, 3, 4, 2, 4, 2}
	//	value := 4
	//	r3 := without(r1, value)
	//	fmt.Println(r1, value)
	//	fmt.Println(r3)
	//}

	//func without(r1 []int, value int) []int {
	//	r2 := []int{}
	//	for j := 0; j < len(r1); j++ {
	//		if  r1[j] != value {
	//			r2 = append(r2, r1[j])
	//		}
	//	}
	//		return r2
	//}

	// Intersection
	// a1 := []int{2, 1, 6, 5}
	//	a2 := []int{2, 3, 5}
	//	a3 := intersection(a1, a2)
	//	fmt.Println(a1)
	//	fmt.Println(a2)
	//	fmt.Println(a3)
	//}
	//	func intersection(a1 []int, a2 []int) []int{
	//		a3 := []int{}
	//			for i := 0; i < len(a1); i++ {
	//				for j := 0; j < len(a2); j++ {
	//					if a1[i] == a2[j] {
	//						a3 = append(a3, a1[i] )
	//					}
	//				}
	//			}
	//
	//		return a3
	//}

	//	UNIQ
	// u0 := []int{2, 5, 2, 5, 2, 7, 8, 9, 12}
	//	u2 := uniq(u0)
	//	fmt.Println(u0)
	//	fmt.Println(u2)
	//}
	//func uniq(u0 []int)  []int {
	//	u1 := []int{}

	//	for i := 0; i < len(u0); i++ {
	//		for j := 0; j < len(u0); j++ {
	//			if u0[i] == u0[j] {
	//				fmt.Println("repeated element")
	//			} else {
	//				if u0[i] != u0[j] {
	//					u1 = append(u1, u0[i])

	//					}
	//			}
	//		}
	//	}

	//	FILL
	// arr := []int{4, 6, 8, 10}
	//	value := 12
	//	startIndex := 1
	//	endIndex := 2
	//	t := fill(arr, value, startIndex, endIndex)
	//	fmt.Println(arr)
	//	fmt.Println(t)
	//}

	//func fill(arr []int, value, startIndex, endIndex int) ([]int) {
	//	for i := startIndex; i <= endIndex; i++{
	//		arr[i] = value
	//	}
	//	return arr
	//}

	//indexOf
	// 	a := []int{1, 2, 1, 2}
	//	value := 2
	//	b := indexOf(a, value)
	//	fmt.Println(a)
	//	fmt.Println(b)

	//func indexOf(a []int, value int) (ind int){
	//	for i:=0; i<len(a); i++{
	//		if a[i]==value{
	//			ind = i
	//			fmt.Println("Index", i)
	//			break
	//		}
	//	}
	//	return ind
	//}

	//	Reverse

	//	c := []int{1, 2, 3}
	//	c1 := reverse(c)
	//	fmt.Println(c)
	//	fmt.Println(c1)
	//}
	//	func reverse(c []int) ([]int) {
	//		r := []int{}
	//			for i := len(c)-1; i >= 0; i-- {
	//			r =append(r, c[i])
	//	}
	//			return r
	//	}

	//	FILL
	//	func main() {
	//		arr := []int{4, 6, 8, 10}
	//		value := 12
	//		startIndex := 1
	//		endIndex := 2
	//		t := fill(arr, value, startIndex, endIndex)
	//		fmt.Println(arr)
	//		fmt.Println(t)
	//	}
	//	func fill(arr []int, value, startIndex, endIndex int) ([]int) {
	//		for i := startIndex; i <= endIndex; i++{
	//			arr[i] = value
	//		}
	//		return arr
	//	}

	//func main() {
	//	fmt.Println("Hello, playground")
	//
	//	arr := []int{4, 6, 8, 10}
	//	value := 12
	//	startIndex := 1
	//	endIndex := 2
	//	t := fill(arr, value, startIndex, endIndex)
	//	fmt.Println(arr)
	//	fmt.Println(t)
	//}

	//func fill(arr []int, value, startIndex, endIndex int) ([]int) {
	//	for i := startIndex; i <= endIndex; i++{
	//		arr[i] = value
	//	}
	//	return arr
	//}

	//Variadic reverse
	//	s:=variadicReverse(2, 3, 5, 5, 3, 6, 11, 11, 87, 46)
	//	fmt.Println(s)
	//}

	//func variadicReverse(nums...int) []int {
	//	fmt.Print(nums, " ")
	//	r := []int{}
	//	for i := len(nums)-1; i >= 0; i-- {
	//	r =append(r, nums[i])
	//	}
	//				return r
	//	}

	//Шахматная доска
	//	height:= 6
	//	width:= 5
	//	char:="*"
	//	c := chessdesk(height, width, char)
	//	fmt.Println(c)
	//}
	//func chessdesk(height, width int, char string) (str string) {
	//	str = ""
	//	for i := 0; i < height; i++ {
	//		//add whitespace before even  row
	//		if i % 2 == 0 {
	//			str += " "
	//		}
	//		for j := 0; j < width; j++ {
	//			str += " " + char + " "
	//	}
	//	str += "\n"

	//	}

	//return

	//}

	//*Конверты

	//	e1 := Envelope{1555.04, 10066.18}
	//	e2 := Envelope{1.25, 8.03}

	//	s:=compareEnvelopes(e1, e2)
	//	fmt.Println(s)
	//fmt.Println("area", e1.square(), e2.square())

	//Сортировка треугольников

