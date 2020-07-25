package retry

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDoIdentityRequests(t *testing.T) {
	n := 3
	expectedResult := "task"

	ctx := context.TODO()
	res, err := Do(
		n,
		// TODO: test that functions were called 1 time
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithIntervalEqualsTo0IdentityRequests(t *testing.T) {
	n := 3
	intervalMs := 0
	expectedResult := "task"

	ctx := context.TODO()
	res, err := DoWithInterval(
		n,
		intervalMs,
		// TODO: test that functions were called 1 time
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithIntervalIdentityRequests(t *testing.T) {
	n := 3
	intervalMs := 1000
	expectedResult := "task"

	ctx := context.TODO()
	res, err := DoWithInterval(
		n,
		intervalMs,
		// TODO: test that functions were called 1 time
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests(t *testing.T) {
	n := 2
	// expected return value is interface{}
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		n,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoWithIntervalIdentityErrorRequests(t *testing.T) {
	// expected return value is interface{}
	n := 2
	intervalMs := 1000
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	// TODO: check that interval works
	res, err := DoWithInterval(
		n,
		intervalMs,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoIncorrectParamNEqualsTo0(t *testing.T) {
	n := 0
	expectedError := "Incorrect n value"

	defer func() {
		err := recover()

		if err == nil {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	Do(
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithIntervalIncorrectParamNEqualsTo0(t *testing.T) {
	n := 0
	intervalMs := 1000
	expectedError := "Incorrect n value"

	defer func() {
		err := recover()

		if err == nil {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithInterval(
		n,
		intervalMs,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoIncorrectParamNEqualsToNeg(t *testing.T) {
	n := -1
	intervalMs := 1000
	expectedError := "Incorrect n value"

	defer func() {
		err := recover()

		if err == nil {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithInterval(
		n,
		intervalMs,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithIntervalIncorrectParamIntervalMsEqualsToNeg(t *testing.T) {
	n := 5
	intervalMs := -1
	expectedError := "Incorrect intervalMs value"

	defer func() {
		err := recover()

		if err == nil {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("retry.DoWithInterval failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithInterval(
		n,
		intervalMs,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func identifyOK(_ context.Context, str string) (string, error) {
	return str, nil
}

func identifyError(_ context.Context, err error) (string, error) {
	return "", err
}
