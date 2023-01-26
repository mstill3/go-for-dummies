package main

import "fmt"

type List[T any] struct {
	head, tail *element[T]
}

func (lst List[T]) String() string {
	return fmt.Sprintf("%v", lst.getAll())
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	    } else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	    }
}

func (lst *List[T]) getAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func myGenerics() {
	// generics aka type parameters
	nums := List[int]{}
	nums.push(10)
	nums.push(54)
	fmt.Println(nums)
}
