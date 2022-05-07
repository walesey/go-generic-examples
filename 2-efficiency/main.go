package main

import (
	"fmt"
	"strconv"
)

// define filter and unique funtions that reuse the original slice memory
// This saves on memory allocations and garbage collection

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

	// [a b c 1 2 3 4 5 1 2 3 4 5]
	// [a b c 1 2 3 4 5]
	// [5 4 3 1 2]
}
