package model

import "time"

type Status struct {
	URL        int
	Clock      time.Time
	StatusCode int
}
