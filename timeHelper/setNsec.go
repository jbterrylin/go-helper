package timehelper

import "time"

func SetNsec(req time.Time, nsec int) (result time.Time) {
	result = time.Date(req.Year(), req.Month(), req.Day(), req.Hour(), req.Minute(), req.Second(), nsec, req.Location())
	return
}
