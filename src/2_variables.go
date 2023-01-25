package main

import (
	"fmt"
	"math"
)

func myVariables() {
	// variables
	var a = "abc"
	var a2 string = "abc"
	var b, c int = 1, 2
	d := "goose"
	fmt.Println(a, a2, b, c, d)

	// constants
	const pi = 3.14
	fmt.Println(math.Sin(pi))
}
