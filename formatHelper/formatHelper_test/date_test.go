package formathelper_test

import (
	"testing"
	"time"

	formathelper "github.com/jbterrylin/go-helper/formatHelper"
)

func TestFormatDateTime(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	t1 := time.Date(2022, time.September, 1, 10, 0, 0, 0, loc)

	expected := "2022-09-01 10:00:00"
	result := formathelper.FormatDateTime(t1)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatDate(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	t1 := time.Date(2022, time.September, 1, 10, 0, 0, 0, loc)

	expected := "2022-09-01"
	result := formathelper.FormatDate(t1)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFormatTime(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	t1 := time.Date(2022, time.September, 1, 10, 0, 0, 0, loc)

	expected := "10:00:00"
	result := formathelper.FormatTime(t1)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
