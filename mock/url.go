package mock

import (
	"github.com/elahe-dastan/HTTP_monitoring/model"
)

type URL struct {
	Urls map[string]int
}

func (u URL) Insert(url model.URL) error {
	u.Urls[url.URL] = url.Period

	return nil
}

func (u URL) GetTable() ([]model.URL, error) {
	models := make([]model.URL, 0)

	for k, v := range u.Urls {
		m := model.URL{
			ID:       0,
			UserID:   0,
			URL:      k,
			Period:   v,
			Statuses: nil,
		}

		models = append(models, m)
	}

	return models, nil
}
