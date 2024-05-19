package timehelper

import "time"

func SetSec(req time.Time, sec int) (result time.Time) {
	result = time.Date(req.Year(), req.Month(), req.Day(), req.Hour(), req.Minute(), sec, req.Nanosecond(), req.Location())
	return
}
