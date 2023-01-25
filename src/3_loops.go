package main

import "fmt"

func myLoops() {
	// for loop
	for i := 1; i <= 8; i++ {
		fmt.Println(i)
	}

	// for each index, letter (rune) in range of string
	// chars are runes in Go
	for _, aRune := range "Hello" {
		fmt.Printf("%c\n", aRune)
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
