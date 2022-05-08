package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Here we can use constraints to define generic functions that make use of += (Sum) and > (Max) operators

type Number interface {
	constraints.Float | constraints.Integer
}

func Sum[T Number](list ...T) T {
	var sum T
	for _, elem := range list {
		sum += elem
	}
	return sum
}

func Max[T constraints.Ordered](args ...T) T {
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	floats := []float32{5.2, -1.3, 0.7, -3.8, 2.6}

	iSum := Sum(ints...)
	fSum := Sum(floats...)
	mixedSum := Sum[float32](1.5, 1, 3, 3.8, -3)

	iMax := Max(ints...)
	fMax := Max(floats...)
	mixedMax := Max(1.5, 3.8)
	stringMax := Max("orange", "pear", "apple")

	fmt.Println(iSum)
	fmt.Println(fSum)
	fmt.Println(mixedSum)

	fmt.Println(iMax)
	fmt.Println(fMax)
	fmt.Println(mixedMax)
	fmt.Println(stringMax)

	// output:

	// 15
	// 3.3999999
	// 6.3

	// 5
	// 2.6
	// 3.8
	// pear
}
