package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)

var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "View All questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI(baseURL)
		if err != nil {
			return err
		}
		err = quizCLI.Questions()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(questionsCmd)
	questionsCmd.Flags().StringVar(&baseURL, "host", "http://localhost:8080", "Base URL")
}
