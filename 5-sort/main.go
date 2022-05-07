package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sort"
)

/*
	Convenient wrapper for:
	sort.Float64s
	sort.Ints
	sort.String
*/

// sort.Ordered
func SortOrdered[T constraints.Ordered](list []T) {
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
}

func main() {
	strs := []string{"c", "b", "d", "f", "g", "e", "a"}
	ints := []int64{3, 1, 8, 5, 4, 9, 0, 2, 6, 7}
	floats := []float32{5.2, -1.3, 0.7, -3.8, 2.6}

	SortOrdered(strs)
	SortOrdered(ints)
	SortOrdered(floats)

	fmt.Println(strs)
	fmt.Println(ints)
	fmt.Println(floats)

	// [a b c d e f g]
	// [0 1 2 3 4 5 6 7 8 9]
	// [-3.8 -1.3 0.7 2.6 5.2]
}
