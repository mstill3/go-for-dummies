package main

import (
	"fmt"
	"time"
)

func showRunes(from string) {
	for index, char := range from {
		fmt.Printf("[%v] %c\n", index, char)
	}
}
func myGoroutine() {
	showRunes("hello")
	// execute concurrently
	go showRunes("sup")
	go showRunes("goodbye")

	// wait for async functions to finish
	time.Sleep(time.Second)
    	fmt.Println("done")
}
