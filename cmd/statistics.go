package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Compare your result",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI(baseURL)
		if err != nil {
			return err
		}
		err = quizCLI.Statistics(userId)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVar(&userId, "user-id", "", "User ID")
	statsCmd.Flags().StringVar(&baseURL, "host", "http://localhost:8080", "Base URL")
	statsCmd.MarkFlagRequired("user-id")
}
