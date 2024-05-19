package timehelper

import "time"

func SetHour(req time.Time, hour int) (result time.Time) {
	result = time.Date(req.Year(), req.Month(), req.Day(), hour, req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}
