package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI("http://localhost:8080")
		if err != nil {
			return err
		}
		err = quizCLI.Register("ebose")
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(registerCmd)
}
