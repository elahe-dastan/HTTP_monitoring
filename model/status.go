package model

import "time"

type Status struct {
	url URL
	clock time.Time
	statusCode int
}