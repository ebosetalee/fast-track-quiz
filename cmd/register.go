package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var userId string
var baseURL string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Registers a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI(baseURL)
		if err != nil {
			return err
		}
		err = quizCLI.Register(userId)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVar(&userId, "user-id", "", "User ID")
	registerCmd.Flags().StringVar(&baseURL, "host", "http://localhost:8080", "Base URL")
	registerCmd.MarkFlagRequired("user-id")
}
