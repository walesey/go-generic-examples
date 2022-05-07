package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func Map[T any, P any](items []T, fn func(T) P) []P {
	results := make([]P, 0, len(items))
	for _, elem := range items {
		results = append(results, fn(elem))
	}
	return results
}

func IntToString(i int) string { return fmt.Sprint(i) }

func main() {
	ints := []int{1, 2, 3, 4, 5}
	strs := Map(ints, IntToString)

	fmt.Println(strs[0] + strs[1] + strs[2] + strs[3] + strs[4])
	fmt.Println(ints[0] + ints[1] + ints[2] + ints[3] + ints[4])

	// output:
	// 12345
	// 15
}
