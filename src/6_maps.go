package main

import "fmt"

func myMap() {
	grades := make(map[string]float32)
	grades["Matt"] = 100.0
	grades["Riley"] = 95.5
	grades["Kaylu"] = 89.3

	fmt.Println(grades)
	// for key, value := range map
	for name, grade := range grades {
		fmt.Println(name, "has grade", grade)
	}
}
