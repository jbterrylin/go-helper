package timehelper_test

import (
	"testing"
	"time"

	timehelper "github.com/jbterrylin/go-helper/timeHelper"
)

func TestIsSameDay(t *testing.T) {
	loc1, _ := time.LoadLocation("America/New_York")
	// loc2, _ := time.LoadLocation("Asia/Tokyo")

	// Test case 1: Same day, different times, ignore time zone
	t1 := time.Date(2022, time.September, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, time.September, 1, 23, 0, 0, 0, loc1)
	if !timehelper.IsSameDay(t1, t2, true) {
		t.Errorf("Expected true, got false")
	}

	if timehelper.IsSameDay(t1, t2, false) {
		t.Errorf("Expected true, got false")
	}
}
