package server

import (
	"HTTP_monitoring/config"
	"HTTP_monitoring/memory"
	"HTTP_monitoring/service"
	"HTTP_monitoring/store"
	"database/sql"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *sql.DB, cfg config.JWT, r redis.Conn) {
	c := cobra.Command{
		Use:   "server",
		Short: "Run server to serve the requests",
		Run: func(cmd *cobra.Command, args []string) {
			URL := store.NewURL(d)
			api := service.API{
				User:   store.NewUser(d),
				URL:    URL,
				Config: cfg,
			}

			du, err := cmd.Flags().GetInt("duration")
			if err != nil {
				log.Fatal(err)
			}
			s := service.Server{
				URL:      URL,
				Status:   store.NewStatus(d),
				Duration: du,
				Redis:     memory.NewStatus(r),
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
