package timehelper

import "time"

func SetDay(req time.Time, day int) (result time.Time) {
	result = time.Date(req.Year(), req.Month(), day, req.Hour(), req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}
