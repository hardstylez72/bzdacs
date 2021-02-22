package tests

import (
	"time"
)

func ParseTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}
	}
	return t
}
