package manage

import (
	"fmt"
	"time"
)

func createdAgo(createdAt string) string {
	layoutISO := "2006-01-02T15:04:05Z"
	converted, _ := time.Parse(layoutISO, createdAt)
	now := time.Now()
	s := now.Sub(converted).Seconds()
	if s < 60 {
		return fmt.Sprintf("%.0f seconds", s)
	} else if s < 3600 {
		return fmt.Sprintf("%.0f minutes", s/60)
	} else if s < 86400 {
		return fmt.Sprintf("%.0f hours", s/3600)
	} else if s > 86400 {
		return fmt.Sprintf("%.0f days", s/86400)
	}

	return ""
}

func timeDiff(createdAt string, resolvedAt string) int {
	layoutISO := "2006-01-02T15:04:05Z"
	ca, _ := time.Parse(layoutISO, createdAt)
	ra, _ := time.Parse(layoutISO, resolvedAt)
	diff := ra.Sub(ca)
	return int(diff.Seconds())
}

func toDate(d string) time.Time {
	layoutISO := "2006-01-02T15:04:05Z"
	c, _ := time.Parse(layoutISO, d)
	return c
}
