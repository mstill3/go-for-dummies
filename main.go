package main

import "fmt"

func main() {
	num := 1
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	switch num {
	case 1:
		myPrinting()
	case 2:
		myVariables()
	case 3:
		myLoops()
	case 4:
		myIf()
	}
}
