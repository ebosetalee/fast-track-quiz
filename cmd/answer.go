package main

import (
	"github.com/ebosetalee/quiz"
	"github.com/spf13/cobra"
)


var ans string

var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer the question",
	RunE: func(cmd *cobra.Command, args []string) error {
		quizCLI, err := quiz.NewCLI(baseURL)
		if err != nil {
			return err
		}
		err = quizCLI.Answer(userId, ans)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(answerCmd)
	answerCmd.Flags().StringVar(&userId, "user-id", "", "User ID")
	answerCmd.Flags().StringVar(&baseURL, "host", "http://localhost:8080", "Base URL")
	answerCmd.Flags().StringVarP(&ans, "option", "a", "", "Answer to the question")
	answerCmd.MarkFlagRequired("user-id")
	answerCmd.MarkFlagRequired("option")
}
