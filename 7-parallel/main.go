package main

import (
	"fmt"
	"os"
	"sync"
)

// A generic implementation of a parallel task runner using go routines, mutex and waitgroup to run
// an arbitrary list of tasks in parallel, return the results in the order the task are specified.
// If an error occurs in any of the tasks, it will be returned along with the index number of the failed task.
// Each task can return any type, so long as they all return the same type.

func Parallel[T any](tasks ...func() (T, error)) ([]T, int, error) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	results := make([]T, len(tasks))
	errs := make([]error, len(tasks))

	wg.Add(len(tasks))

	for index := range tasks {
		go func(i int) {
			defer wg.Done()
			result, err := tasks[i]()
			mu.Lock()
			results[i] = result
			errs[i] = err
			mu.Unlock()
		}(index)
	}

	wg.Wait()

	for i, err := range errs {
		if err != nil {
			return results, i, err
		}
	}

	return results, -1, nil
}

func main() {
	files, errNb, err := Parallel[[]byte](
		func() ([]byte, error) { return os.ReadFile("./main.go") },
		func() ([]byte, error) { return os.ReadFile("./README.md") },
		func() ([]byte, error) { return os.ReadFile("./go.mod") },
		func() ([]byte, error) { return os.ReadFile("./go.sum") },
	)

	fmt.Println(len(files[0]), len(files[1]), len(files[2]), len(files[3]))
	fmt.Println(err)
	fmt.Println(errNb)
	// output:
	// 536 89 91 207
	// <nil>
	// -1
}
