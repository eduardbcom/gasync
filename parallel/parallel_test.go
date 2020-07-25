package parallel

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDoIdentityRequests(t *testing.T) {
	// expected return value is interface{}
	expectedResult := []interface{}{"task1", "task2"}

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResult, res)
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

func TestDoWithLimitIdentityIncorrectLimitEqualsTo0(t *testing.T) {
	limit := 0
	expectedError := "Incorrect limit value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
	)
}

func TestDoWithLimitIdentityIncorrectLimitIsNeg(t *testing.T) {
	limit := -1
	expectedError := "Incorrect limit value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
	)
}

func TestDoIdentityErrorRequests1(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedError, res)
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

func TestDoIdentityErrorRequests2(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
	)
	if res != nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedError, res)
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

func TestDoIdentityErrorRequests3(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("parallel.Do failed. Expected %v, actual: %v", expectedError, res)
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
