package migrate

import (
	"HTTP_monitoring/store"
	"database/sql"

	"github.com/spf13/cobra"
)

//nolint: gofumpt
func Register(root *cobra.Command, d *sql.DB) {
	c := cobra.Command{
		Use:   "migrate",
		Short: "Manages database, creates and fills tables if don't exist",
		Run: func(cmd *cobra.Command, args []string) {
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
