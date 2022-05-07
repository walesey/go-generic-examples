package main

import (
	"fmt"
	"sync"
	"time"
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
	results, _, err := Parallel[string](
		func() (string, error) {
			time.Sleep(2 * time.Second)
			return "1", nil
		},
		func() (string, error) {
			time.Sleep(1 * time.Second)
			return "2", nil
		},
		func() (string, error) {
			time.Sleep(4 * time.Second)
			return "3", nil
		},
		func() (string, error) {
			time.Sleep(3 * time.Second)
			return "4", nil
		},
	)

	fmt.Println(results)
	fmt.Println(err)
	// output
	// [1 2 3 4]
	// <nil>

	results, errNb, err := Parallel[string](
		func() (string, error) {
			return "1", nil
		},
		func() (string, error) {
			return "failed", fmt.Errorf("failed")
		},
		func() (string, error) {
			return "3", nil
		},
	)

	fmt.Println(results)
	fmt.Println(err)
	fmt.Println(errNb)
	// output
	// [1 failed 3]
	// failed
	// 1

}
