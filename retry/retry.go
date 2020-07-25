package retry

import (
	"time"
)

// Do attempts to get a successful response from function `f`
// no more than `times` times before returning an error.
//
// If the function call to `f` is successful,
// this function returns right away with the result value.
//
// Otherwise, it returns and error of the last attempt.
func Do(times int, f func() (interface{}, error)) (interface{}, error) {
	if times <= 0 {
		panic("Incorrect n value")
	}

	var lastError error

	for i := 0; i < times; i++ {
		res, err := f()
		if err == nil {
			return res, nil
		}

		lastError = err
	}

	return nil, lastError
}

// DoWithInterval is the same as Do
// but sleep for `intervalMs` milliseconds before next call.
func DoWithInterval(times int, intervalMs int, f func() (interface{}, error)) (interface{}, error) {
	if times <= 0 {
		panic("Incorrect n value")
	}

	if intervalMs < 0 {
		panic("Incorrect intervalMs value")
	}

	if intervalMs == 0 {
		return Do(times, f)
	}

	var lastError error

	for i := 0; i < times; i++ {
		res, err := f()
		if err == nil {
			return res, nil
		}

		lastError = err

		// dont sleep on the last call
		if i < times -1 {
			time.Sleep(time.Duration(intervalMs) * time.Millisecond)
		}
	}

	return nil, lastError
}