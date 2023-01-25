package main

import "fmt"

func count100(onHalfComplete func()) {
	for index := 1; index <= 10; index++ {
		fmt.Println(index)
		if index == 5 {
			onHalfComplete()
		}
	}
}

func myClosures() {
	// closures are functions that references variables from outside its body.
	// The function may access and assign to the referenced variables
	secretNum := 42
	count100(func() {
		fmt.Println("Half way done! ---", secretNum)
	})
}
