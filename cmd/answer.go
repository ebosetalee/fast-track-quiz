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
		quizCLI, err := quiz.NewCLI("http://localhost:8080")
		if err != nil {
			return err
		}
		err = quizCLI.Answer("ebose", ans)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	cliCmd.AddCommand(answerCmd)
	answerCmd.Flags().StringVarP(&ans, "option", "a", "", "Answer to the question")
	answerCmd.MarkFlagRequired("option")
}
