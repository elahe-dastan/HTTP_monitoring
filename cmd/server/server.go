package server

import (
	"HTTP_monitoring/config"
	"HTTP_monitoring/service"
	"HTTP_monitoring/store"
	"database/sql"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *sql.DB, cfg config.JWT) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				URL := store.NewURL(d)
				api := service.API{
					User: store.NewUser(d),
					URL:  URL,
					Config:cfg,
				}
				s := service.Server{
					URl:    URL,
					Status: store.NewStatus(d),
				}
				go s.Run()
				api.Run()
			},
		},
	)
}