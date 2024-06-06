package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetSec(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Sec      int
		Expected time.Time
	}{
		{
			Name:     "Set to new second",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Sec:      45,
			Expected: time.Date(2023, 1, 1, 12, 0, 45, 0, location),
		},
		{
			Name:     "Keep same second",
			Req:      time.Date(2023, 1, 1, 12, 0, 30, 0, location),
			Sec:      30,
			Expected: time.Date(2023, 1, 1, 12, 0, 30, 0, location),
		},
		{
			Name:     "Change to zero second",
			Req:      time.Date(2023, 1, 1, 12, 0, 59, 123456789, location),
			Sec:      0,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 123456789, location),
		},
		{
			Name:     "Change second with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Sec:      50,
			Expected: time.Date(2023, 1, 1, 8, 30, 50, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetSec(testCase.Req, testCase.Sec)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
