package model

import "time"

type Status struct {
	URL        int	`redis:"url"`
	Clock      time.Time	`redis:"clock"`
	StatusCode int	`redis:"status"`
}
