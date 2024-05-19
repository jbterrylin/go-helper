package timehelper

import "time"

func SetDate(req time.Time, year int, month time.Month, day int) (result time.Time) {
	result = time.Date(year, month, day, req.Hour(), req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}
