package main

import (
	"fmt"
	"regexp"
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

		"four": 4,
		"4":    4,
	}

	keys := Keys(m)

	greaterThan2 := Filter(m, func(k string, v int) bool { return v > 2 })
	noThree := Filter(m, func(k string, v int) bool { return k != "three" && v != 3 })
	alphaNumericKey := Filter(m, func(k string, v int) bool { return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(k) })

	fmt.Println(keys)
	fmt.Println(greaterThan2)
	fmt.Println(noThree)
	fmt.Println(alphaNumericKey)

	// output
	// [two 2 three 3 four 4 one 1]
	// map[3:3 4:4 four:4 three:3]
	// map[1:1 2:2 4:4 four:4 one:1 two:2]
	// map[four:4 one:1 three:3 two:2]

}
