package timehelper

import (
	"fmt"
	"time"
)

func IsSameDay(t1, t2 time.Time, ignoreTimeZone bool) bool {
	if !ignoreTimeZone {
		t1 = t1.UTC()
		t2 = t2.UTC()
	}
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	fmt.Println(t1, y1, m1, d1)
	fmt.Println(t2, y2, m2, d2)

	return y1 == y2 && m1 == m2 && d1 == d2
}
