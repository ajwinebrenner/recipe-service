package models

import "time"

type Chef struct {
	ID        int
	FirstName string
	LastName  string
	DOB       time.Time
}
