package parallel

import (
	"sync"
)

// Do runs functions in parallel.
//
// If one of them returns an error, the final result is a (nil, Error).
// The amount of function calls in that case is undefined.
//
// Otherwise, it's an array of results.
func Do(funcs ...func() (interface{}, error)) ([]interface{}, error) {
	return DoWithLimit(len(funcs), funcs...)
}

// DoWithLimit runs maximum `limit` amount of functions in parallel.
//
// If one of them returns an error, the final result is a (nil, Error).
// The amount of function calls in that case is undefined.
//
// Otherwise, it's an array of results.
func DoWithLimit(limit int, funcs ...func() (interface{}, error)) ([]interface{}, error) {
	if len(funcs) == 0 {
		return make([]interface{}, 0), nil
	}

	if limit <= 0 {
		panic("Incorrect limit value")
	}

	results := make([]interface{}, len(funcs))

	var wg sync.WaitGroup

	ch := make(chan bool, limit)
	errorCh := make(chan error, limit)

	for i, f := range funcs {
		// limit amount of parallel functions
		ch <- true

		if len(errorCh) != 0 {
			break
		}

		wg.Add(1)

		go func(i int, f func() (interface{}, error)) {
			defer wg.Done()
			defer func() { <-ch }()

			res, err := f()
			if err != nil {
				errorCh <- err

				return
			}

			results[i] = res
		}(i, f)
	}

	wg.Wait()

	if len(errorCh) != 0 {
		return nil, <-errorCh
	}

	return results, nil
}
