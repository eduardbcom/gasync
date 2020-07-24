package tryEach

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDoIdentityRequests(t *testing.T) {
	// expected return value is interface{}
	expectedResult := "task1"

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that functions were called only once
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests1(t *testing.T) {
	// expected return value is interface{}
	expectedResult := "task1"

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that both function was called
		func() (interface{}, error) { return identifyOK(ctx, "task1") },
		func() (interface{}, error) { return identifyError(ctx, errors.New("some error here")) },
	)
	if err != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests2(t *testing.T) {
	// expected return value is interface{}
	expectedResult := "task2"

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that second function wasn't called
		func() (interface{}, error) { return identifyError(ctx, errors.New("some error here")) },
		func() (interface{}, error) { return identifyOK(ctx, "task2") },
	)
	if err != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests3(t *testing.T) {
	// expected return value is interface{}
	var expectedResult interface{} = nil
	expectedError := errors.New("some second error here")

	ctx := context.TODO()
	res, err := Do(
		// TODO: test that second function wasn't called
		func() (interface{}, error) { return identifyError(ctx, errors.New("some error here")) },
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("series.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func identifyOK(ctx context.Context, str string) (string, error) {
	return str, nil
}

func identifyError(ctx context.Context, err error) (string, error) {
	return "", err
}
