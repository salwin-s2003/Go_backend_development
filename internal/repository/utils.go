package repository

import (
	"time"
)

func ParseDate(dob string) (time.Time, error) {
	// Expecting format YYYY-MM-DD
	return time.Parse("2006-01-02", dob)
}
