package tools

import "time"

func FormatISODate(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
