package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetDate(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Year     int
		Month    time.Month
		Day      int
		Expected time.Time
	}{
		{
			Name:     "Set to new date",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2022,
			Month:    time.February,
			Day:      20,
			Expected: time.Date(2022, 2, 20, 12, 0, 0, 0, location),
		},
		{
			Name:     "Keep same date",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2023,
			Month:    time.January,
			Day:      1,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change year",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2024,
			Month:    time.January,
			Day:      1,
			Expected: time.Date(2024, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change month",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2023,
			Month:    time.February,
			Day:      1,
			Expected: time.Date(2023, 2, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change day",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2023,
			Month:    time.January,
			Day:      31,
			Expected: time.Date(2023, 1, 31, 12, 0, 0, 0, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetDate(testCase.Req, testCase.Year, testCase.Month, testCase.Day)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
