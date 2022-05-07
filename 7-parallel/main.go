package main

import (
	"fmt"
	"time"
)

func Parallel[T any](tasks ...func() (T, error)) ([]T, []error) {
	chans := make([]chan T, len(tasks))
	for i := range chans {
		chans[i] = make(chan T)
	}
	eChans := make([]chan error, len(tasks))
	for i := range eChans {
		eChans[i] = make(chan error)
	}

	for index := range tasks {
		go func(i int) {
			result, err := tasks[i]()
			chans[i] <- result
			eChans[i] <- err
		}(index)
	}

	results := make([]T, 0, len(tasks))
	for _, c := range chans {
		results = append(results, <-c)
	}

	errs := make([]error, 0, len(tasks))
	for _, c := range eChans {
		errs = append(errs, <-c)
	}

	return results, errs
}

func main() {
	results, errs := Parallel[string](
		func() (string, error) {
			time.Sleep(2 * time.Second)
			return "1", nil
		},
		func() (string, error) {
			time.Sleep(1 * time.Second)
			return "failed", fmt.Errorf("failed")
		},
		func() (string, error) {
			time.Sleep(6 * time.Second)
			return "2", nil
		},
		func() (string, error) {
			time.Sleep(3 * time.Second)
			return "3", nil
		},
	)

	fmt.Println(results)
	fmt.Println(errs)

	// [1 failed 2 3]
	// [<nil> failed <nil> <nil>]
}
