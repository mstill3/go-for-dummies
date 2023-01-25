package main

import "fmt"

func plus(num1 int, num2 int) int {
	return num1 + num2
}

// can omit parameter types when several same
func plusPlus(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

// multiple return values
func getResponse() (int, string) {
	return 400, "error"
}

// variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func myFunctions() {
	num1 := 2
	num2 := 5
	simpleSum := plus(num1, num2)
	fmt.Println(num1, "+", num2, "=", simpleSum)
	fmt.Println(num1, "+", num2, "+", simpleSum, "=", plusPlus(num1, num2, simpleSum))

	respCode, respMessage := getResponse()
	fmt.Println(respCode, "---", respMessage)

	fmt.Println(sum(2, 6, 4, 3, 2, 5, 3))

	// anonymous/lambda function
	mySum := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	fmt.Println(num1, "+", num2, "=", mySum(num1, num2))
}
