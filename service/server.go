package service

import (
	"time"

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
		if counter == s.Threshold {
			statuses := s.Redis.Flush()
			for i := 0; i < len(statuses); i++ {
				s.Status.Insert(statuses[i])
				//; err != nil {
				//	fmt.Println(err)
				//}
			}
			counter = 0
		}

		//rows := s.URL.GetTable()
		//
		////nolint: bodyclose
		//for rows.Next() {
		//	var url model.URL
		//
		//	if err := rows.Scan(&url.ID, &url.UserID, &url.URL, &url.Period); err != nil {
		//		fmt.Println(err)
		//	}
		//
		//	if counter % url.Period != 0 {
		//		continue
		//	}
		//
		//	//nolint: noctx
		//	resp, err := http.Get(url.URL)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//
		//	var status model.Status
		//	status.URL = url.ID
		//	status.Clock = time.Now().Format(time.Stamp)
		//	status.StatusCode = resp.StatusCode
		//
		//	s.Redis.Insert(status)
		//}
	}
}
