package main

import "fmt"

// NOTE: uses Person struct in `12_structs.go`
// similar to Java toString()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

type IPAddr [4]byte
func (address IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", address[0], address[1], address[2], address[3])
}

func myStringer() {
	// Stringer is a very common interface
	// A Stringer is a type that can describe itself as a string
	person1 := Person{name: "Matt", age:25}
	fmt.Println(person1)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
