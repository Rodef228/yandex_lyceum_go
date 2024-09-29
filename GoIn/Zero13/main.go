package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	diff := now.Sub(pastTime)

	years := int(diff.Hours() / 24 / 365)
	months := int(diff.Hours() / 24 / 30)
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes())
	seconds := int(diff.Seconds())

	switch {
	case years > 0:
		return fmt.Sprintf("%d year%s ago", years, pluralize(years))
	case months > 0:
		return fmt.Sprintf("%d month%s ago", months, pluralize(months))
	case days > 0:
		return fmt.Sprintf("%d day%s ago", days, pluralize(days))
	case hours > 0:
		return fmt.Sprintf("%d hour%s ago", hours, pluralize(hours))
	case minutes > 0:
		return fmt.Sprintf("%d minute%s ago", minutes, pluralize(minutes))
	default:
		return fmt.Sprintf("%d second%s ago", seconds, pluralize(seconds))
	}
}

func pluralize(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

