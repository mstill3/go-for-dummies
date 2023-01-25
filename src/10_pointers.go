package main

import "fmt"

func add2Immutable(num int) int {
	return num + 2
}

func add2Mutable(num *int) {
	*num += 2
}

func myPointers() {
	num := 9
	pointerToNum := &num
	fmt.Println(pointerToNum, ": ", num)

	newNum := add2Immutable(num)
	fmt.Println(num, "add 2 =", newNum)
	add2Mutable(&num)
	fmt.Println("new num value", num)
}
