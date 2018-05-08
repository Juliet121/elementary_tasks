package main

import (
	"math"
	"fmt"
)

type Triangle struct {
	Vertices string
	sideA float64
	sideB float64
	sideC float64
}

func checkinputdata() []Triangle {
	triangles := []Triangle{}
	for i := 0; i < 3; i++ {
		t1 := Triangle{}
		fmt.Println("Please enter the symbolic name of vertices of the ", i+1, " triangle:    ")
		fmt.Scanf("%s", &t1.Vertices)
		fmt.Println("Please enter side A of the ", i+1, " triangle:     ")
		fmt.Scan(&t1.sideA)
		fmt.Println("Please enter side B of the ", i+1, "triangle:     ")
		fmt.Scan(&t1.sideB)
		fmt.Println("Please enter side C of the ", i+1, " triangle:     ")
		fmt.Scan(&t1.sideC)

		if t1.sideA+t1.sideB < t1.sideC || t1.sideA+t1.sideC < t1.sideB || t1.sideB+t1.sideC < t1.sideA ||
			t1.sideA < 0 || t1.sideB < 0 || t1.sideC < 0 {
			fmt.Println("User entered incorrect input data (negative digit) or it's not possible to build the" +
				"  triangle with such side lenght. Please re-start the app")
			break
		}else{
			triangles = append(triangles, t1)
		}
	}
	return triangles
}

func (c *Triangle) halfperimeter() (p float64) {
	p = (c.sideC+c.sideB+c.sideA)/2
	return p
}

func (c *Triangle) area() (a float64) {
	a = math.Sqrt(c.halfperimeter()*(c.halfperimeter()-c.sideA)*(c.halfperimeter()-c.sideB)*(c.halfperimeter()-c.sideC))
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
	fmt.Println("Title of triangles: ", ComparisonofTriangles(checkinputdata()))
}

