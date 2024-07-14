package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var port int64 

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Starts the http server",
	Run: func(cmd *cobra.Command, args []string) {
		quiz.Main(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().Int64Var(&port, "port", 8080, "port")
}
