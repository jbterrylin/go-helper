package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetYear(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Year     int
		Expected time.Time
	}{
		{
			Name:     "Set to new year",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2025,
			Expected: time.Date(2025, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Keep same year",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2023,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to past year",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2000,
			Expected: time.Date(2000, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to future year",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Year:     2050,
			Expected: time.Date(2050, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change year with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Year:     2024,
			Expected: time.Date(2024, 1, 1, 8, 30, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetYear(testCase.Req, testCase.Year)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
