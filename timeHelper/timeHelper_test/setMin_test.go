package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetMin(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Min      int
		Expected time.Time
	}{
		{
			Name:     "Set to new minute",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Min:      45,
			Expected: time.Date(2023, 1, 1, 12, 45, 0, 0, location),
		},
		{
			Name:     "Keep same minute",
			Req:      time.Date(2023, 1, 1, 12, 30, 0, 0, location),
			Min:      30,
			Expected: time.Date(2023, 1, 1, 12, 30, 0, 0, location),
		},
		{
			Name:     "Change to beginning of hour",
			Req:      time.Date(2023, 1, 1, 12, 30, 0, 0, location),
			Min:      0,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to end of hour",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Min:      59,
			Expected: time.Date(2023, 1, 1, 12, 59, 0, 0, location),
		},
		{
			Name:     "Change minute with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Min:      15,
			Expected: time.Date(2023, 1, 1, 8, 15, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetMin(testCase.Req, testCase.Min)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
