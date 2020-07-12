package model

import "time"

type Status struct {
	Url        int
	Clock      time.Time
	StatusCode int
}
