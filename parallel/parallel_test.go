package parallel

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type expectedResult struct {
	res []interface{}
	err error
}

func TestDoIdentityRequests(t *testing.T) {
	expectedResults := []expectedResult{
		expectedResult{[]interface{}{}, nil},
		expectedResult{[]interface{}{"task1"}, nil},
		expectedResult{[]interface{}{"task1", "task2"}, nil},
		expectedResult{[]interface{}{"task1", "task2", "task3"}, nil},
		expectedResult{nil, errors.New("some error here 0")},
		expectedResult{nil, errors.New("some error here 1")},
		expectedResult{nil, errors.New("some error here 2")},
		expectedResult{nil, errors.New("some error here 3")},
	}

	ctx := context.TODO()

	testCases := [][]func() (interface{}, error) {
		[]func() (interface{}, error) {},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
			func() (interface{}, error) { return identifyOK(ctx, "task2") },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
			func() (interface{}, error) { return identifyOK(ctx, "task2") },
			func() (interface{}, error) { return identifyOK(ctx, "task3") },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyError(ctx, errors.New("some error here 0")) },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyError(ctx, errors.New("some error here 1")) },
			func() (interface{}, error) { return identifyOK(ctx, "task2") },
			func() (interface{}, error) { return identifyOK(ctx, "task3") },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
			func() (interface{}, error) { return identifyError(ctx, errors.New("some error here 2")) },
			func() (interface{}, error) { return identifyOK(ctx, "task3") },
		},
		[]func() (interface{}, error) {
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
			func() (interface{}, error) { return identifyOK(ctx, "task2") },
			func() (interface{}, error) { return identifyError(ctx, errors.New("some error here 3")) },
		},
	}

	for i, testCase := range testCases {
		res, err := Do(
			// TODO: test that functions were called only once
			testCase...
		)
		if !reflect.DeepEqual(expectedResults[i].err, err) {
			t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResults[i], err)
		}

		if !reflect.DeepEqual(expectedResults[i].res, res) {
			t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResults[i], res)
		}
	}
}

func TestDoWithLimitIdentityIncorrectLimit(t *testing.T) {
	testFunc := func(limit int) {
		expectedError := "Incorrect limit value"

		defer func() {
			err := recover()

			if !reflect.DeepEqual(expectedError, err) {
				t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
			}
		}()

		ctx := context.TODO()
		DoWithLimit(
			limit,
			func() (interface{}, error) { return identifyOK(ctx, "task1") },
		)

		panic("Should not be here")
	}

	limits := []int{
		-1,
		0,
	}

	for _, limit := range limits {
		testFunc(limit)
	}
}

func TestDoWithLimitIdentityRequests1(t *testing.T) {
	limit := 1
	// expected return value is interface{}
	expectedResult := []interface{}{"task1", "task2"}

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithLimitIdentityRequests2(t *testing.T) {
	limit := 2
	// expected return value is interface{}
	expectedResult := []interface{}{"task1", "task2"}

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithLimitIdentityRequests3(t *testing.T) {
	limit := 10
	// expected return value is interface{}
	expectedResult := []interface{}{"task1", "task2"}

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithLimitIdentityErrorRequests1(t *testing.T) {
	limit := 1
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoWithLimitIdentityErrorRequests2(t *testing.T) {
	limit := 1
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
	)
	if res != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoWithLimitIdentityErrorRequests3(t *testing.T) {
	limit := 1
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoWithLimitIdentityErrorRequests4(t *testing.T) {
	limit := 2
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoWithLimitIdentityErrorRequests5(t *testing.T) {
	limit := 100
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.DoWithLimit failed. Expected %v, actual: %v", expectedError, res)
	}
}

func identifyOK(_ context.Context, str string) (string, error) {
	return str, nil
}

func identifyError(_ context.Context, err error) (string, error) {
	return "", err
}
