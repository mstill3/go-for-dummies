package main

import "fmt"

type any interface{}

func do(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func myTypes() {
	// A type assertion provides access to an interface value's underlying concrete value
	var myVar interface{} = "hello"
	myVarString, okStr := myVar.(string)
	myVarInt, okInt := myVar.(int)
	myVarFloat, okFloat := myVar.(float64)
	
	fmt.Println("Is String:", myVarString, okStr)
	fmt.Println("Is Int:", myVarInt, okInt)
	fmt.Println("Is Float:", myVarFloat, okFloat)

	do(3)
	do("hi")
	do(5.6)
}
