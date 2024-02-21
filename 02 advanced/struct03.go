package main

import "fmt"

type Circle struct {
	Radius float64
}

func (cir Circle) area() float64 {
	area := 3.14 * cir.Radius * cir.Radius
	return area
}

func main() {
	var circle Circle
	circle.Radius = 5
	area := circle.area()
	fmt.Println("面积:", area)
}
