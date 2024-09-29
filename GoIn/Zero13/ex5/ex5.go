package ex5

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	nextDay := start.AddDate(0, 0, 1)

	for nextDay.Weekday() == time.Saturday || nextDay.Weekday() == time.Sunday {
		nextDay = nextDay.AddDate(0, 0, 1)
	}

	return nextDay
}
