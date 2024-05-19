package timehelper

import "time"

func SetTime(req time.Time, hour, min, sec int, nsec ...int) (result time.Time) {
	nanoSecond := req.Nanosecond()
	if len(nsec) > 0 {
		nanoSecond = nsec[0]
	}
	result = time.Date(req.Year(), req.Month(), req.Day(), hour, min, sec, nanoSecond, req.Location())
	return
}
