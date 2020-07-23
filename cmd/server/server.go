package server

import (
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/service"
	"github.com/elahe-dastan/HTTP_monitoring/store"
	"github.com/elahe-dastan/HTTP_monitoring/store/status"
	"github.com/nats-io/go-nats"
	"gorm.io/gorm"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *gorm.DB, jwt config.JWT, r redis.Conn, threshold int, n *nats.Conn, natsConfig config.Nats) {
	c := cobra.Command{
		Use:   "server",
		Short: "Run server to serve the requests",
		Run: func(cmd *cobra.Command, args []string) {
			URL := store.NewURL(d)
			api := service.API{
				User:   store.NewUser(d),
				URL:    URL,
				Config: jwt,
			}

			du, err := cmd.Flags().GetInt("duration")
			if err != nil {
				log.Fatal(err)
			}
			s := service.Server{
				URL:       URL,
				Status:    status.NewSQLStatus(d),
				Duration:  du,
				Redis:     status.NewRedisStatus(r),
				Threshold: threshold,
				NatsConn:  n,
				NatsCfg:   natsConfig,
			}
			go s.Run()
			api.Run()
		},
	}

	c.Flags().IntP("duration", "d", 1,
		"every d minutes the status of the urls will be checked")

	root.AddCommand(
		&c,
	)
}
