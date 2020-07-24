package tryEach

// Do runs each func in series but stops whenever any of the functions were successful.
//
// If one of the tasks were successful,
// the return value will be the result of the successful func.
//
// If all tasks fail
// the return value will be and error and result (if any) of the final attempt.
func Do(funcs ...func() (interface{}, error)) (interface{}, error) {
	var lastError error = nil

	for _, f := range funcs {
		res, err := f()
		if err == nil {
			return res, nil
		}

		lastError = err
	}

	return nil, lastError
}
