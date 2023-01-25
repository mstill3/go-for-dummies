package main

import "fmt"

func myMap() {
	grades := make(map[string]float32)
	//grades := map[string]float32{"Billy": 74.3}
	grades["Matt"] = 100.0
	grades["Riley"] = 95.5
	grades["Kaylu"] = 89.3

	delete(grades, "Kaylu")

	fmt.Println(grades)
	// for key, value := range map
	for name, grade := range grades {
		fmt.Println(name, "has grade", grade)
	}
}
