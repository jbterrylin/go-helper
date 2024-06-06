package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetMonth(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Month    time.Month // Can be time.Month or int
		Expected time.Time
	}{
		{
			Name:     "Set to new month",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    time.February,
			Expected: time.Date(2023, 2, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Keep same month",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    time.January,
			Expected: time.Date(2023, 1, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to December",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    time.December,
			Expected: time.Date(2023, 12, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change month with different time",
			Req:      time.Date(2023, 1, 15, 8, 30, 45, 123456789, location),
			Month:    time.March,
			Expected: time.Date(2023, 3, 15, 8, 30, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetMonth(testCase.Req, testCase.Month)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}

func TestSetMonthByInt(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Month    int
		Expected time.Time
	}{
		{
			Name:     "Set to new month by int",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    2,
			Expected: time.Date(2023, 2, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Keep same month by int",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    1,
			Expected: time.Date(2023, 1, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change to December by int",
			Req:      time.Date(2023, 1, 15, 12, 0, 0, 0, location),
			Month:    12,
			Expected: time.Date(2023, 12, 15, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change month by int with different time",
			Req:      time.Date(2023, 1, 15, 8, 30, 45, 123456789, location),
			Month:    3,
			Expected: time.Date(2023, 3, 15, 8, 30, 45, 123456789, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetMonthByInt(testCase.Req, testCase.Month)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
