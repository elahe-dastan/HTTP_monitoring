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
	NatsConn  *nats.Conn
	NatsCfg   config.Nats
}

func (s *Server) Run() {
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

		for _, u := range urls {
			if counter%u.Period != 0 {
				continue
			}

			s.Publish(u)
		}
	}
}

func (s *Server) Publish(u model.URL) {
	ec, err := nats.NewEncodedConn(s.NatsConn, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	err = ec.Publish(s.NatsCfg.Topic, u)
	if err != nil {
		log.Fatal(err)
	}
}
