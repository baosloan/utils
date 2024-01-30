package timeutil

import (
	"time"
)

const (
	dateOnlyLength = 10
	DateTimeLength = 19
)

// GetDateOnly 获取年月日 格式为: YYYY-MM-DD
func GetDateOnly(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.DateOnly)
}

// GetDateTime 获取年月日时分秒 YYYY-MM-DD HH:MM:SS
func GetDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.DateTime)
}

// GetTimeOnly 获取时分秒 HH:MM:SS
func GetTimeOnly(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.TimeOnly)
}
