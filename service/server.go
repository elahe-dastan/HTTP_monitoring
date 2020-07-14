package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/elahe-dastan/HTTP_monitoring/store"
)

type Server struct {
	URL      store.SQLURL
	Status   store.SQLStatus
	Duration int
	Redis    store.RedisStatus
	Threshold int
}

func (s *Server) Run() {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

	//nolint: sqlclosecheck
	for {
		<-ticker.C

		counter++
		if counter == 6 {
			models := s.Redis.Flush()
			for i := 0; i < len(models); i++ {
				if err := s.Status.Insert(models[i]); err != nil {
					fmt.Println(err)
				}
			}
			counter = 1
		}

		rows := s.URL.GetTable()

		//nolint: bodyclose
		for rows.Next() {
			var url model.URL

			if err := rows.Scan(&url.ID, &url.UserID, &url.URL, &url.Period); err != nil {
				fmt.Println(err)
			}

			if counter % url.Period != 0 {
				continue
			}

			//nolint: noctx
			resp, err := http.Get(url.URL)
			if err != nil {
				fmt.Println(err)
			}

			var status model.Status
			status.URL = url.ID
			status.Clock = time.Now().String()
			status.StatusCode = resp.StatusCode

			s.Redis.Insert(status)
		}
	}
}
