package mock

import "github.com/elahe-dastan/HTTP_monitoring/model"

type Status struct {
}

func (s *Status) Insert(status model.Status) error {
	return nil
}
