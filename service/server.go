package service

import (
	"fmt"
	"log"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/elahe-dastan/HTTP_monitoring/store"
	"github.com/elahe-dastan/HTTP_monitoring/store/status"
	"github.com/nats-io/go-nats"
)

type Server struct {
	URL       store.SQLURL
	Status    status.SQLStatus
	Duration  int
	Redis     status.RedisStatus
	Threshold int
	Nats      *nats.Conn
}

func (s *Server) Run(cfg config.Nats) {
	ticker := time.NewTicker(time.Duration(s.Duration) * time.Minute)
	counter := 0

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
			fmt.Println(counter)
			if counter%u.Period != 0 {
				continue
			}

			s.Publish(u, cfg)
		}
	}
}

func (s *Server) Publish(u model.URL, c config.Nats) {
	ec, err := nats.NewEncodedConn(s.Nats, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = ec.Publish(c.Topic, u)
	if err != nil {
		log.Fatal(err)
	}
}