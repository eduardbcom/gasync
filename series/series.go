package series

// Do runs the functions in the `funcs` collection in series.
// Each one running once the previous function has completed.
//
// If any functions in the series return an error,
// no more functions are run, and `series.Do` is immediately returned with the value of the error.
//
// Otherwise, function returns an array of results when tasks have completed.
func Do(funcs ...func() (interface{}, error)) ([]interface{}, error) {
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
