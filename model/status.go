package model

import "time"

type Status struct {
	URL        URL	`redis:"url"`
	Clock      time.Time	`redis:"clock"`
	StatusCode int	`redis:"status"`
}
