package timehelper

import "time"

func IsBetween(startDate, endDate, targetDate time.Time) bool {
	if targetDate.Equal(startDate) || targetDate.Equal(endDate) {
		return true
	}
	return targetDate.After(startDate) && targetDate.Before(endDate)
}

func IsBetweenOnly(startDate, endDate, targetDate time.Time) bool {
	return targetDate.After(startDate) && targetDate.Before(endDate)
}
