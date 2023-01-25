package main

import (
	"fmt"
	"unicode/utf8"
)

func myStrings() {
	message := "Hello World"
	fmt.Println(message)
	fmt.Println("length: ", utf8.RuneCountInString(message))

	for index, runeValue := range message {
		var theRune rune = runeValue
		fmt.Println(index, theRune)
	}
}
