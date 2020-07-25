package times

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestDoIdentityRequests(t *testing.T) {
	n := 3
	expectedResult := []interface{}{"task", "task", "task"}

	ctx := context.TODO()
	res, err := Do(
		n,
		// TODO: test that functions were called N times
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoWithLimitIdentityRequests(t *testing.T) {
	limit := 3
	n := 3
	expectedResult := []interface{}{"task", "task", "task"}

	ctx := context.TODO()
	res, err := DoWithLimit(
		limit,
		n,
		// TODO: test that functions were called N times
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoSeriesLimitIdentityRequests(t *testing.T) {
	n := 3
	expectedResult := []interface{}{"task", "task", "task"}

	ctx := context.TODO()
	res, err := DoSeries(
		n,
		// TODO: test that functions were called N times
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
	if err != nil {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, err)
	}

	if !reflect.DeepEqual(expectedResult, res) {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, res)
	}
}

func TestDoIdentityErrorRequests2(t *testing.T) {
	n := 3
	var expectedResult interface{}
	expectedError := errors.New("some error here")

	ctx := context.TODO()
	res, err := Do(
		n,
		func() (interface{}, error) { return identifyError(ctx, expectedError) },
	)
	if res != nil {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedResult, res)
	}

	if err == nil {
		t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, res)
	}
}

func TestDoIncorrectParamNEqualsTo0(t *testing.T) {
	n := 0
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	Do(
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithLimitIncorrectParamLimitEqualsTo0(t *testing.T) {
	limit := 0
	n := 1
	expectedError := "Incorrect limit value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithLimitIncorrectParamLimitEqualsToNeg(t *testing.T) {
	limit := -1
	n := 1
	expectedError := "Incorrect limit value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithLimitIncorrectParamNEqualsTo0(t *testing.T) {
	limit := 1
	n := 0
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoWithLimitIncorrectParamNEqualsToNeg(t *testing.T) {
	limit := 1
	n := -1
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.DoWithLimit failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoWithLimit(
		limit,
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoSeriesLimitIncorrectParamNEqualsTo0(t *testing.T) {
	n := 0
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.DoSeries failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.DoSeries failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoSeries(
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoIncorrectParamNEqualsToNeg(t *testing.T) {
	n := -1
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	Do(
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func TestDoSeriesIncorrectParamNEqualsToNeg(t *testing.T) {
	n := -1
	expectedError := "Incorrect n value"

	defer func() {
		err := recover();
		
		if err == nil {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}

		if !reflect.DeepEqual(expectedError, err) {
			t.Errorf("times.Do failed. Expected %v, actual: %v", expectedError, err)
		}
	}()

	ctx := context.TODO()
	DoSeries(
		n,
		func() (interface{}, error) { return identifyOK(ctx, "task") },
	)
}

func identifyOK(_ context.Context, str string) (string, error) {
	return str, nil
}

func identifyError(_ context.Context, err error) (string, error) {
	return "", err
}
