package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetHour(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Hour     int
		Expected time.Time
	}{
		{
			Name:     "Set to new hour",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Hour:     15,
			Expected: time.Date(2023, 1, 1, 15, 0, 0, 0, location),
		},
		{
			Name:     "Keep same hour",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Hour:     12,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to midnight",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Hour:     0,
			Expected: time.Date(2023, 1, 1, 0, 0, 0, 0, location),
		},
		{
			Name:     "Change to noon",
			Req:      time.Date(2023, 1, 1, 0, 0, 0, 0, location),
			Hour:     12,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change hour with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Hour:     20,
			Expected: time.Date(2023, 1, 1, 20, 30, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetHour(testCase.Req, testCase.Hour)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
