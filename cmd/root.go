package cmd

import (
	"HTTP_monitoring/cmd/migrate"
	"HTTP_monitoring/cmd/server"
	"HTTP_monitoring/config"
	"HTTP_monitoring/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "monitor",
		Short: "A project that checks the status of URLs",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	exitFailure := 1

	cfg := config.Read()
	d := db.New(cfg.Database)

	migrate.Register(rootCmd, d)
	server.Register(rootCmd, d)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}

