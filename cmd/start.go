package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz or continue from where you left off",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI(baseURL)
		if err != nil {
			return err
		}
		err = quizCLI.Start(userId)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&userId, "user-id", "", "User ID")
	startCmd.Flags().StringVar(&baseURL, "host", "http://localhost:8080", "Base URL")
	startCmd.MarkFlagRequired("user-id")
}
