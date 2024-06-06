package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetTime(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Hour     int
		Min      int
		Sec      int
		Nsec     []int
		Expected time.Time
	}{
		{
			Name:     "Set hour, minute, second",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Hour:     10,
			Min:      30,
			Sec:      45,
			Nsec:     nil,
			Expected: time.Date(2023, 1, 1, 10, 30, 45, 0, location),
		},
		{
			Name:     "Set hour, minute, second, nanosecond",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Hour:     10,
			Min:      30,
			Sec:      45,
			Nsec:     []int{500},
			Expected: time.Date(2023, 1, 1, 10, 30, 45, 500, location),
		},
		{
			Name:     "Keep nanosecond if not provided",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 123456789, location),
			Hour:     10,
			Min:      30,
			Sec:      45,
			Nsec:     nil,
			Expected: time.Date(2023, 1, 1, 10, 30, 45, 123456789, location),
		},
		{
			Name:     "Change hour, keep minute and second",
			Req:      time.Date(2023, 1, 1, 12, 30, 45, 0, location),
			Hour:     15,
			Min:      30,
			Sec:      45,
			Nsec:     nil,
			Expected: time.Date(2023, 1, 1, 15, 30, 45, 0, location),
		},
		{
			Name:     "Change nanosecond",
			Req:      time.Date(2023, 1, 1, 12, 30, 45, 123456789, location),
			Hour:     12,
			Min:      30,
			Sec:      45,
			Nsec:     []int{987654321},
			Expected: time.Date(2023, 1, 1, 12, 30, 45, 987654321, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetTime(testCase.Req, testCase.Hour, testCase.Min, testCase.Sec, testCase.Nsec...)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
