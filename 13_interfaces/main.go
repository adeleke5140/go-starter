package main

import (
	"fmt"
	"math"
)

//Define Interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}


//I understand this as saying:
//for our Circle struct, we are gonna have an area method
//that takes a circle instance and returns the area with a 
//type of float64
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//This one here is for the rectagle method
func (r Rectangle) area() float64 {
	return r.width * r.height
}  

//I wonder what this method is doing
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x:10, y:20, radius: 4}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", rectangle.area())
}