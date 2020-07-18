package migrate

import (
	"github.com/elahe-dastan/HTTP_monitoring/store"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, d *gorm.DB) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
			user := store.NewUser(d)
			user.Create()

			url := store.NewURL(d)
			url.Create()

			status := store.NewSQLStatus(d)
			status.Create()
		},
	}

	root.AddCommand(
		&c,
	)
}
