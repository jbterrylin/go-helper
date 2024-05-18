package formathelper

import "time"

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// FormatDate 格式化时间为 YYYY-MM-DD 格式
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatTime 格式化时间为 HH:MM:SS 格式
func FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}
