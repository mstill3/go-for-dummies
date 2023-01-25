package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// receiver type of Rectangle
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect Rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

func myMethods() {
	myRectangle := Rectangle{width: 20, height: 40}
	fmt.Println("Rectangle:", myRectangle)
	fmt.Println("Area:", myRectangle.area())
	fmt.Println("Perimeter:", myRectangle.perimeter())
}
