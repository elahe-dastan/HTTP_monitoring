package subscriber

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/model"
	"github.com/elahe-dastan/HTTP_monitoring/store/status"
	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, n *nats.Conn, cfg config.Nats, r status.RedisStatus) {
	c := cobra.Command{
		Use: "subscribe",
		Run: func(cmd *cobra.Command, args []string) {
			Subscribe(n, cfg, r)
		},
	}

	root.AddCommand(
		&c,
	)
}

func Subscribe(nc *nats.Conn, cfg config.Nats, r status.RedisStatus) {
	c, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	ch := make(chan model.URL)

	if _, err := c.QueueSubscribe(cfg.Topic, cfg.Queue, func(u model.URL) {
		ch <- u
	}); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		go worker(ch, r)
	}

	select {}
}

//nolint: bodyclose
func worker(ch chan model.URL, r status.RedisStatus) {
	for u := range ch {
		resp, err := http.Get(u.URL)
		if err != nil {
			fmt.Println(err)
		}

		var st model.Status
		st.URLID = u.ID
		st.Clock = time.Now()
		st.StatusCode = resp.StatusCode

		r.Insert(st)
	}
}
