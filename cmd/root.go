package cmd

import (
	"HTTP_monitoring/cmd/migrate"
	"HTTP_monitoring/cmd/server"
	"HTTP_monitoring/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "monitor",
		Short: "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	exitFailure := 1

	cfg := config.Read()

	migrate.Register(rootCmd, cfg)
	server.Register(rootCmd, cfg)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitFailure)
	}
}

