package main

import "fmt"

// NOTE: uses Person struct in `12_structs.go`
// similar to Java toString()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func myStringer() {
	// Stringer is a very common interface
	// A Stringer is a type that can describe itself as a string
	person1 := Person{name: "Matt", age:25}
	fmt.Println(person1)
}
