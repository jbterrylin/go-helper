package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestSetNsec(t *testing.T) {
	location, _ := time.LoadLocation("UTC")

	testCases := []struct {
		Name     string
		Req      time.Time
		Nsec     int
		Expected time.Time
	}{
		{
			Name:     "Set to new nanosecond",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 0, location),
			Nsec:     500,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 500, location),
		},
		{
			Name:     "Keep same nanosecond",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 300, location),
			Nsec:     300,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 300, location),
		},
		{
			Name:     "Change to zero nanosecond",
			Req:      time.Date(2023, 1, 1, 12, 0, 0, 123456789, location),
			Nsec:     0,
			Expected: time.Date(2023, 1, 1, 12, 0, 0, 0, location),
		},
		{
			Name:     "Change nanosecond with different time",
			Req:      time.Date(2023, 1, 1, 8, 30, 45, 123456789, location),
			Nsec:     987654321,
			Expected: time.Date(2023, 1, 1, 8, 30, 45, 987654321, location),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.SetNsec(testCase.Req, testCase.Nsec)
			if !result.Equal(testCase.Expected) {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
