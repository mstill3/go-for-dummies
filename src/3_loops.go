package main

import "fmt"

func myLoops() {
	// for loop
	for i := 1; i <= 8; i++ {
		fmt.Println(i)
	}

	// for each index, letter in range of string
	for _, character := range "Hello" {
		fmt.Printf("%c\n", character)
	}

	// while loop
	running := true
	for running {
		num := 0
		fmt.Println("Enter a number:")
		fmt.Scan(&num)
		if num == 0 {
			running = false
		}
	}

}
