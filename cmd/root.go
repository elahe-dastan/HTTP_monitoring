package cmd

import (
	"fmt"
	"os"

	"github.com/elahe-dastan/HTTP_monitoring/cmd/migrate"
	"github.com/elahe-dastan/HTTP_monitoring/cmd/server"
	"github.com/elahe-dastan/HTTP_monitoring/cmd/subscriber"
	"github.com/elahe-dastan/HTTP_monitoring/config"
	"github.com/elahe-dastan/HTTP_monitoring/db"
	"github.com/elahe-dastan/HTTP_monitoring/load_balancer"
	"github.com/elahe-dastan/HTTP_monitoring/memory"

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
	n := load_balancer.Conn(cfg.Nats)

	migrate.Register(rootCmd, d)
	server.Register(rootCmd, d, cfg.JWT, r, cfg.Redis.Threshold, n, cfg.Nats)
	subscriber.Register(rootCmd, n, cfg.Nats, r)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}
