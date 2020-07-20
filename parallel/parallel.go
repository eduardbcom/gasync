package gasync

// Parallel runs functions in parallel.
// If one of them returns an error, the final result is a (nil, Error).
// Otherwise, it's an array of results
func Parallel(funcs ...func() (interface{}, error)) ([]interface{}, error) {
	results := make([]interface{}, len(funcs))

	for i, f := range funcs {
		res, err := f()
		if err != nil {
			return nil, err
		}

		results[i] = res
	}

	return results, nil
}