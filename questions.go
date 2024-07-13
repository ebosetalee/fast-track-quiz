package quiz

import "sort"

type Question struct {
	index         int
	Text          string   `json:"question"`
	Answers       []answer `json:"answers"`
	correctAnswer string
}

type answer struct {
	Index string `json:"index"`
	Text  string `json:"text"`
}

type QuizStruct struct {
	Questions []Question
}

type UserAnswer struct {
	ChosenAnswer string `json:"answer" valid:"required~please provide an answer"`
}

func (q *QuizStruct) sort() *QuizStruct {
	sort.SliceStable(q.Questions, func(i int, j int) bool {
		return q.Questions[i].index < q.Questions[j].index
	})

	return q
}

func (q *QuizStruct) getQuestion(index int) *Question {
	for i := 0; i < len(q.Questions); i++ {
		if index == q.Questions[i].index {
			return &q.Questions[i]
		}
	}
	return nil
}

func (q *Question) checkAnswer(answer string) string {
	if q.correctAnswer == answer {
		return "Correct"
	}
	return "Wrong"
}

var Quiz = QuizStruct{
	Questions: []Question{
		{
			index: 2,
			Text:  "Which planet has the most moons?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "Saturn",
				},
				{
					Index: "B",
					Text:  "Jupiter",
				},
				{
					Index: "C",
					Text:  "Earth",
				},
				{
					Index: "D",
					Text:  "Uranus",
				},
			},
			correctAnswer: "A",
		},
		{
			index: 1,
			Text:  "How many colors are in the rainbow?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "8",
				},
				{
					Index: "B",
					Text:  "9",
				},
				{
					Index: "C",
					Text:  "7",
				},
				{
					Index: "D",
					Text:  "12",
				},
			},
			correctAnswer: "C",
		},
	},
}
