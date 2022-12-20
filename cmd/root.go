package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "activator",
	Short: "Activatator allows you to manage fonts without leaving the command line.",
	Args:  cobra.NoArgs,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = "[0.1.0]"
}
