package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/elahe-dastan/HTTP_monitoring/store"
	"github.com/elahe-dastan/HTTP_monitoring/store/status"
)

type Server struct {
	URL       store.SQLURL
	Status    status.SQLStatus
	Duration  int
	Redis     status.RedisStatus
	Threshold int
}

func (s *Server) Run() {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

	//nolint: sqlclosecheck
	for {
		<-ticker.C

		counter++
		if counter == s.Threshold {
			statuses := s.Redis.Flush()
			for i := 0; i < len(statuses); i++ {
				if err := s.Status.Insert(statuses[i]); err != nil {
					fmt.Println(err)
				}
			}
			counter = 1
		}

		urls, err := s.URL.GetTable()
		if err != nil {
			log.Fatal(err)
		}

		//nolint: bodyclose
		for _, u := range urls {

			if counter % u.Period != 0 {
				continue
			}

			//nolint: noctx
			resp, err := http.Get(u.URL)
			if err != nil {
				fmt.Println(err)
			}

			var st model.Status
			st.URLID = u.ID
			st.Clock = time.Now()
			st.StatusCode = resp.StatusCode

			s.Redis.Insert(st)
		}
	}
}
