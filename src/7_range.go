package main

import "fmt"

func myRange() {
	nums := []int{3, 6, 4, 5}
	sum := 0
	for _, num := range nums {
		fmt.Print(num, "+")
		sum += num
	}
	fmt.Println("=", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	myMap := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range myMap {
		fmt.Printf("%s -> %s\n", key, val)
	}

	for key := range myMap {
		fmt.Println("key:", key)
	}

	for index, aRune := range "go" {
		fmt.Println(index, aRune)
	}
}
