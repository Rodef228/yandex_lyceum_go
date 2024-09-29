package ex1

import (
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}