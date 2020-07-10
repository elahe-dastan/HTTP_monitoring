package migrate

import (
	"HTTP_monitoring/config"
	"HTTP_monitoring/db"
	"HTTP_monitoring/store"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, cfg config.Config) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			d := db.New(cfg.Database)

			user := store.NewUser(d)
			user.Create()

			url := store.NewURL(d)
			url.Create()

			status := store.NewStatus(d)
			status.Create()
		},
	}

	root.AddCommand(
		&c,
	)
}
