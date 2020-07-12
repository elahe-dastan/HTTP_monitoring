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

	for {
		<-ticker.C
		rows := s.URL.GetTable()
		for rows.Next() {
			var url model.URL

			if err := rows.Scan(&url.Id, &url.UserId, &url.Url); err != nil {
				fmt.Println(err)
			}

			resp, err := http.Get(url.Url)
			if err != nil {
				fmt.Println(err)
			}

			var status model.Status
			status.Url = url.Id
			status.Clock = time.Now()
			status.StatusCode = resp.StatusCode


			if err := s.Status.Insert(status); err != nil {
				fmt.Println(err)
			}

		}
	}
}