package main

import (
	"fmt"
	"encoding/json"
)

func myJson() {
	// to JSON string
	fruits := []string{"apple", "peach", "pear"}
	fruitsBytes, _ := json.Marshal(fruits)
	fruitsJsonStr := string(fruitsBytes)
	fmt.Println(fruitsJsonStr)

	// from JSON string to data structure
	fruitsBytes2 := []byte(fruitsJsonStr)
	var readInFruits []string
	if err := json.Unmarshal(fruitsBytes2, &readInFruits); err != nil {
		panic(err)
	}
	for _, fruit := range readInFruits {
		fmt.Println(fruit)
	}
}
