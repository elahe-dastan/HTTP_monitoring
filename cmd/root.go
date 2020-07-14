package cmd

import (
	"HTTP_monitoring/cmd/migrate"
	"HTTP_monitoring/cmd/server"
	"HTTP_monitoring/config"
	"HTTP_monitoring/db"
	"HTTP_monitoring/memory"
	"HTTP_monitoring/redis"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
//nolint: gofumpt
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "monitor",
		Short: "A project that checks the status of URLs",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	exitFailure := 1

	cfg := config.Read()
	d := db.New(cfg.Database)
	r := memory.New(cfg.Redis)

	migrate.Register(rootCmd, d)
	server.Register(rootCmd, d, cfg.JWT, r)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}
