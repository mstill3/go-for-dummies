package main

import "fmt"

func main() {
	num := 1
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	fmt.Println()

	switch num {
	case 1:
		myPrinting()
	case 2:
		myVariables()
	case 3:
		myLoops()
	case 4:
		myIf()
	case 5:
		myArray()
	case 6:
		myMap()
	case 7:
		myRange()
	case 8:
		myFunctions()
	case 9:
		myClosures()
	case 10:
		myPointers()
	case 11:
		myStrings()
	case 12:
		myStructs()
	case 13:
		myMethods()
	case 14:
		myInterfaces()
	case 15:
		myTypes()
	case 16:
		myStringer()
	case 17:
		myErrors()
	case 18:
		myReaders()
	case 19:
		myGenerics()
	case 20:
		myGoroutine()
	default:
		fmt.Println("ERROR - Unrecognized option")
	}
}
