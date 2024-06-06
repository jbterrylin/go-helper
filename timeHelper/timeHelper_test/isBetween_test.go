package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestIsBetween(t *testing.T) {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	testCases := []struct {
		Name       string
		StartDate  time.Time
		EndDate    time.Time
		TargetDate time.Time
		Expected   bool
	}{
		{
			Name:       "Target date is start date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: startDate,
			Expected:   true,
		},
		{
			Name:       "Target date is end date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: endDate,
			Expected:   true,
		},
		{
			Name:       "Target date is between",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC),
			Expected:   true,
		},
		{
			Name:       "Target date is before start date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC),
			Expected:   false,
		},
		{
			Name:       "Target date is after end date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.IsBetween(testCase.StartDate, testCase.EndDate, testCase.TargetDate)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}

func TestIsBetweenOnly(t *testing.T) {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	testCases := []struct {
		Name       string
		StartDate  time.Time
		EndDate    time.Time
		TargetDate time.Time
		Expected   bool
	}{
		{
			Name:       "Target date is start date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: startDate,
			Expected:   false,
		},
		{
			Name:       "Target date is end date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: endDate,
			Expected:   false,
		},
		{
			Name:       "Target date is between",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC),
			Expected:   true,
		},
		{
			Name:       "Target date is before start date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC),
			Expected:   false,
		},
		{
			Name:       "Target date is after end date",
			StartDate:  startDate,
			EndDate:    endDate,
			TargetDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := timehelper.IsBetweenOnly(testCase.StartDate, testCase.EndDate, testCase.TargetDate)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
