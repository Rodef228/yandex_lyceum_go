package ex2

import (
	"time"
)

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}
