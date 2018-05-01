package main

import (
	"math"
	"fmt"
)

type Triangle struct {
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

