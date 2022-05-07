package main

import (
	"fmt"
	"strconv"
)

// If you wanted to filter a very large list you won't want to create an entire copy of the list.
// Define filter and unique funtions that reuse the original slice memory.
// These functions are passed a list by reference to manipulate the original slice via pointer.
// This saves on memory allocations and garbage collection.
// Note: the order is not preserved, as this allows us to reuse the slice memory for max performance.

func Filter[T any](list *[]T, fn func(T) bool) {
	for i := 0; i < len(*list); {
		elem := (*list)[i]
		if fn(elem) {
			i++
		} else {
			*list = RemoveItem(*list, i)
		}
	}
}

func Unique[T comparable](list *[]T) {
	seen := map[T]bool{}
	for i := 0; i < len(*list); {
		elem := (*list)[i]
		if _, ok := seen[elem]; ok {
			*list = RemoveItem(*list, i)
		} else {
			seen[elem] = true
			i++
		}
	}
}

func RemoveItem[T any](list []T, i int) []T {
	list[i] = list[len(list)-1]
	return list[:len(list)-1]
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	strs := []string{"a", "b", "c", "1", "2", "3", "4", "5", "1", "2", "3", "4", "5"}
	fmt.Println(strs)

	Unique(&strs)
	fmt.Println(strs)

	Filter(&strs, IsInt)
	fmt.Println(strs)

	// output
	// [a b c 1 2 3 4 5 1 2 3 4 5]
	// [a b c 1 2 3 4 5]
	// [5 4 3 1 2]
}
