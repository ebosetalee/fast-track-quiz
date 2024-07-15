package quiz

import (
	"sort"
	"sync"
)

type User struct {
	Id    string
	Score int
}

type Database struct {
	Users map[string]User
	mu    *sync.Mutex
}

type Response struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

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

type UserRequest struct {
	Id string `json:"Id" valid:"required~please provide an Id"`
}

func NewDatabase() *Database {
	return &Database{
		Users: make(map[string]User, 10),
		mu:    &sync.Mutex{},
	}
}

func (d *Database) register(user User) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Users[user.Id] = user
}

func (d *Database) addScore(userId string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	score := d.Users[userId].Score + 1

	d.Users[userId] = User{
		Id:    userId,
		Score: score,
	}
}

func (d *Database) sortByScore() []User {
	allUsers := make([]User, 0)

	for _, user := range d.Users {
		allUsers = append(allUsers, user)
	}

	sort.SliceStable(allUsers, func(i int, j int) bool {
		return allUsers[i].Score > allUsers[j].Score
	})

	return allUsers
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
			index: 4,
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
		{
			index: 1,
			Text:  "Who was the Ancient Greek God of the Sun?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "Athena",
				},
				{
					Index: "B",
					Text:  "Apollo",
				},
				{
					Index: "C",
					Text:  "Artemis",
				},
				{
					Index: "D",
					Text:  "Ares",
				},
			},
			correctAnswer: "B",
		},
		{
			index: 5,
			Text:  "Which of these shapes has 7 sides?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "Hexagon",
				},
				{
					Index: "B",
					Text:  "Pentagon",
				},
				{
					Index: "C",
					Text:  "Octagon",
				},
				{
					Index: "D",
					Text:  "Heptagon",
				},
			},
			correctAnswer: "D",
		},
		{
			index: 3,
			Text:  `What city is known as "The Eternal City"?`,
			Answers: []answer{
				{
					Index: "A",
					Text:  "Athens",
				},
				{
					Index: "B",
					Text:  "Rome",
				},
				{
					Index: "C",
					Text:  "Imperial City",
				},
				{
					Index: "D",
					Text:  "Jerusalem",
				},
			},
			correctAnswer: "B",
		}, 
		{
			index: 6,
			Text:  "Which of these countries is the most densely populated?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "Ontario",
				},
				{
					Index: "B",
					Text:  "Singapore",
				},
				{
					Index: "C",
					Text:  "Monaco",
				},
				{
					Index: "D",
					Text:  "Hong Kong",
				},
			},
			correctAnswer: "C",
		},
		{
			index: 8,
			Text:  "I'll Make a Man Out of You' is a song from which movie?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "Pocahontas",
				},
				{
					Index: "B",
					Text:  "Mulan",
				},
				{
					Index: "C",
					Text:  "Brave",
				},
				{
					Index: "D",
					Text:  "Hong Kong",
				},
			},
			correctAnswer: "B",
		},
		{
			index: 7,
			Text:  "What country is Cognac from?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "France",
				},
				{
					Index: "B",
					Text:  "Spain",
				},
				{
					Index: "C",
					Text:  "Russia",
				},
				{
					Index: "D",
					Text:  "Italy",
				},
			},
			correctAnswer: "A",
		},
		{
			index: 10,
			Text:  "Where did sushi originate?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "China",
				},
				{
					Index: "B",
					Text:  "Korea",
				},
				{
					Index: "C",
					Text:  "Japan",
				},
				{
					Index: "D",
					Text:  "Singapore",
				},
			},
			correctAnswer: "A",
		},
		{
			index: 9,
			Text:  "Kratos is the main character of what video game series?",
			Answers: []answer{
				{
					Index: "A",
					Text:  "BloodBorne",
				},
				{
					Index: "B",
					Text:  "Uncharted",
				},
				{
					Index: "C",
					Text:  "Elden Ring",
				},
				{
					Index: "D",
					Text:  "God Of War",
				},
			},
			correctAnswer: "D",
		},
	},
}
