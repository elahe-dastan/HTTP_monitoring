package server

import (
	"HTTP_monitoring/config"
	"HTTP_monitoring/db"
	"HTTP_monitoring/service"
	"HTTP_monitoring/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				d := db.New(cfg.Database)
				api := service.API{
					User: store.NewUser(d),
				}
				api.Run()
			},
		},
	)
}