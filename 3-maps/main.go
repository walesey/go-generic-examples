package main

import (
	"fmt"
)

func Keys[K comparable, V any](m map[K]V) []K {
	results := make([]K, 0, len(m))
	for k := range m {
		results = append(results, k)
	}
	return results
}

func Filter[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	results := map[K]V{}
	for k, v := range m {
		if fn(k, v) {
			results[k] = v
		}
	}
	return results
}

func main() {
	m := map[string]int{
		"one": 1,
		"1":   1,

		"two": 2,
		"2":   2,

		"three": 3,
		"3":     3,
	}

	keys := Keys(m)

	noThree := Filter(m, func(k string, v int) bool { return k != "three" && v != 3 })

	fmt.Println(keys)
	fmt.Println(noThree)

	// [one 1 two 2 three 3]
	// map[1:1 2:2 one:1 two:2]
}
