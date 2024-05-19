package timehelper

import "time"

func SetMonth(req time.Time, month time.Month) (result time.Time) {
	result = time.Date(req.Year(), month, req.Day(), req.Hour(), req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}

func SetMonthByInt(req time.Time, month int) (result time.Time) {
	result = time.Date(req.Year(), time.Month(month), req.Day(), req.Hour(), req.Minute(), req.Second(), req.Nanosecond(), req.Location())
	return
}
