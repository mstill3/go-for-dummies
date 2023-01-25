package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle Circle) perimeter() float64 {
	return math.Pi * 2 * circle.radius
}

type Geometry interface {
	area() float64
	perimeter() float64
}

func measure(geometry Geometry) {
	fmt.Printf("Shape: %+v\n", geometry)
	fmt.Println("Area:", geometry.area())
	fmt.Println("Perimeter:", geometry.perimeter())
}

func myInterfaces() {
	myRectangle := Rectangle{width: 30.0, height: 50.0}
	myCircle := Circle{radius: 45.5}
	measure(myRectangle)
	measure(myCircle)
}
