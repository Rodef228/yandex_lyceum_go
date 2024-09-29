package ex3

import (
	"time"
)

func ParseStringToTime(dateString, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}
