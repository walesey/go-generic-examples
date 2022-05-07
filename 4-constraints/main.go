package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Sum[T Number](list []T) T {
	var sum T
	for _, elem := range list {
		sum += elem
	}
	return sum
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	floats := []float32{5.2, -1.3, 0.7, -3.8, 2.6}

	iSum := Sum(ints)
	fSum := Sum(floats)

	fmt.Println(iSum)
	fmt.Println(fSum)

	// output:
	// 15
	// 3.3999999
}
