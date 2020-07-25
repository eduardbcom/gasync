package times

import (
	"github.com/eduardbcom/gasync/parallel"
	"github.com/eduardbcom/gasync/series"
)

// Do runs function N times in parallel.
//
// If function returns an error, the final result is a (nil, Error).
// The amount of function calls in that case is undefined.
//
// Otherwise, it's an array of results with length N.
func Do(n int, f func() (interface{}, error)) ([]interface{}, error) {
	if n <= 0 {
		panic("Incorrect n value")
	}

	funcs := make([]func() (interface{}, error), n)
	for i := 0; i < n; i++ {
		funcs[i] = f
	}

	return parallel.Do(funcs...)
}

// DoWithLimit is the same as Do
// but runs a maximum of limit async operations at a time.
func DoWithLimit(limit int, n int, f func() (interface{}, error)) ([]interface{}, error) {
	if limit <= 0 {
		panic("Incorrect limit value")
	}

	if n <= 0 {
		panic("Incorrect n value")
	}

	funcs := make([]func() (interface{}, error), n)
	for i := 0; i < n; i++ {
		funcs[i] = f
	}

	return parallel.DoWithLimit(limit, funcs...)
}

// DoSeries is the same as Do
// but runs only a single function at a time.
func DoSeries(n int, f func() (interface{}, error)) ([]interface{}, error) {
	if n <= 0 {
		panic("Incorrect n value")
	}

	funcs := make([]func() (interface{}, error), n)
	for i := 0; i < n; i++ {
		funcs[i] = f
	}

	return series.Do(funcs...)
}
