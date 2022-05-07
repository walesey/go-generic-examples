package main

import (
	"fmt"
	"strconv"
)

func Filter[T any](list []T, fn func(T) bool) []T {
	results := []T{}
	for _, item := range list {
		if fn(item) {
			results = append(results, item)
		}
	}
	return results
}

func Unique[T comparable](list []T) []T {
	results := make([]T, 0, len(list))
	seen := map[T]bool{}
	for _, elem := range list {
		if _, ok := seen[elem]; !ok {
			seen[elem] = true
			results = append(results, elem)
		}
	}
	return results
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	strs := []string{"a", "b", "c", "1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}

	uniqStrs := Unique(strs)

	intStrs := Filter(uniqStrs, IsInt)

	fmt.Println(strs)
	fmt.Println(uniqStrs)
	fmt.Println(intStrs)

	// output
	// [a b c 1 2 3 4 5 1 2 3 4 5]
	// [a b c 1 2 3 4 5]
	// [1 2 3 4 5]
}
