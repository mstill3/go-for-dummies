package main

import "fmt"

type Person struct {
	name string
	age  int
}

// embedded struct
type Employee struct {
	id int
	//  An embedding looks like a field without a name
	Person
}

// toString Stringer for Employee
func (employee Employee) String() string {
	return fmt.Sprintf("[%v] %v (age %v)\n", employee.id, employee.name, employee.age)
}

// You can safely return a pointer to local variable as a local variable will survive the scope of the function
func createPerson(name string, age int) *Person {
	newPerson := Person{name: name, age: age}
	return &newPerson
}

func myStructs() {
	p1 := Person{name: "Matt", age: 25}
	fmt.Println(p1)

	p1.age = 27
	fmt.Println(p1)

	// You can also use dots with struct pointers - the pointers are automatically de-referenced
	p2 := createPerson("Riley", 6)
	fmt.Println(p2.name)

	e1 := Employee{id: 6, Person: p1}
	fmt.Println(e1)
}
