package timehelper

import "time"

func SetYear(req time.Time, year int) (result time.Time) {
	result = time.Date(year, req.Month(), req.Day(), req.Hour(), req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}
