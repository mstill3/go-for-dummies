package main

import (
	"fmt"
	"time"
)

func myIf() {
	// if
	num := 7
	if num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is off")
	}

	// switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
