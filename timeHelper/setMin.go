package timehelper

import "time"

func SetMin(req time.Time, min int) (result time.Time) {
	result = time.Date(req.Year(), req.Month(), req.Day(), req.Hour(), min, req.Second(), req.Nanosecond(), req.Location())
	return
}
