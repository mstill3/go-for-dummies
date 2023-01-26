package main

import "fmt"

func showRunes(from string) {
	for index, char := range from {
		fmt.Printf("[%v] %c\n", index, char)
	}
}
func myGoroutine() {
	showRunes("hello")
	go showRunes("sup?")
	// go showRunes("goodbye")
}
