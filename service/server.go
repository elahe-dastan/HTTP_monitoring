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
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Second)

	//nolint: sqlclosecheck
	for {
		<-ticker.C

		rows := s.URL.GetTable()

		//nolint: bodyclose
		for rows.Next() {
			var url model.URL

			if err := rows.Scan(&url.ID, &url.UserID, &url.URL); err != nil {
				fmt.Println(err)
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
