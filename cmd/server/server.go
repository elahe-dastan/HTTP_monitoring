package server

import (
	"database/sql"
	"log"

	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/service"
	"github.com/elahe-dastan/HTTP_monitoring/store"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *sql.DB, jwt config.JWT, r redis.Conn, threshold int) {
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
				URL:      URL,
				Status:   store.NewSQLStatus(d),
				Duration: du,
				Redis:     store.NewRedisStatus(r),
				Threshold: threshold,
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
