package main

import (
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Print the cli usage, use the sub commands instead",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Parent().Usage()
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
