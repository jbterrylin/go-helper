package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetDay(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Day      int
		Expected time.Time
	}{
		{
			Name:     "Set to new day",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Day:      15,
			Expected: time.Date(2023, 1, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Keep same day",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Day:      1,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to last day of month",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Day:      31,
			Expected: time.Date(2023, 1, 31, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to first day of month",
			Req:      time.Date(2023, 1, 31, 12, 0, 0, 0, location),
			Day:      1,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change day with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Day:      20,
			Expected: time.Date(2023, 1, 20, 8, 30, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetDay(testCase.Req, testCase.Day)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
