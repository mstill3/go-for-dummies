package main

import "fmt"

func myArray() {
	// arrays
	array := [3]int{3, 6, 9}
	fmt.Println(array)

	// slices
	slice := []int{2, 4, 6}
	fmt.Println(slice)

	slice = append(slice, 8)
	fmt.Println(slice)

	slice = slice[0:2]
	fmt.Println(slice)
}
