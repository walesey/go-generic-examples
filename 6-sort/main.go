package main

import (
	"fmt"
	"sort"
)

// We can even define a generic sort function that takes it's own 'less' function

func Sort[T any](list []T, less func(T, T) bool) {
	sort.Slice(list, func(i, j int) bool { return less(list[i], list[j]) })
}

type Person struct {
	name string
	age  int
}

func ByAge(p1, p2 Person) bool {
	return p1.age < p2.age
}

func ByName(p1, p2 Person) bool {
	return p1.name < p2.name
}

func main() {
	people := []Person{
		{name: "John", age: 34},
		{name: "Sally", age: 23},
		{name: "Bob", age: 47},
	}

	Sort(people, ByAge)
	fmt.Println(people)

	Sort(people, ByName)
	fmt.Println(people)

	// output
	// [{Sally 23} {John 34} {Bob 47}]
	// [{Bob 47} {John 34} {Sally 23}]
}
