package balancer

import (
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/nats-io/go-nats"
)

func New(n config.Nats) *nats.Conn {
	nc, err := nats.Connect(n.Host)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}
