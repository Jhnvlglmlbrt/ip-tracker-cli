package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP-tracker CLi app",
		Long:  `IP-tracker CLi app`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
