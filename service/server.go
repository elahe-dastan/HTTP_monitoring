package service

import (
	"HTTP_monitoring/model"
	"HTTP_monitoring/store"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	URL      store.SQLURL
	Status   store.SQLStatus
	Duration int
}

func (s *Server) Run() {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

	//nolint: sqlclosecheck
	for {
		<-ticker.C

		counter++
		if counter == 101 {
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
			status.Clock = time.Now()
			status.StatusCode = resp.StatusCode

			if err := s.Status.Insert(status); err != nil {
				fmt.Println(err)
			}
		}
	}
}
