package model

type Status struct {
	URL        int	`redis:"url"`
	Clock      string	`redis:"clock"`
	StatusCode int	`redis:"status"`
}
