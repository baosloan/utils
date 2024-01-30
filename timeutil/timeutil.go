package timeutil

import (
	"errors"
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

// ParseInLocationShanghai 将日期字符串转换为time.Time
func ParseInLocationShanghai(timeStr string) (t time.Time, err error) {
	if len(timeStr) != dateOnlyLength && len(timeStr) != DateTimeLength {
		return t, errors.New("timeStr format is invalid")
	}
	//上海时区
	loc, err := time.LoadLocation(LocationShanghai)
	if err != nil {
		return
	}
	//解析时间字符串并指定时区
	if len(timeStr) == dateOnlyLength {
		t, err = time.ParseInLocation(time.DateOnly, timeStr, loc)
	} else if len(timeStr) == DateTimeLength {
		t, err = time.ParseInLocation(time.DateTime, timeStr, loc)
	}
	return
}
