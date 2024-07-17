# QUIZ

This simple quiz features 10 diverse questions, each with multiple-choice options. Sharpen your wits, choose the correct answer from the provided alternatives, and see how your score compares to others. 

### Preferred Stack:
Backend - Golang
Database - Just in-memory, so no database 

### Preferred Components: 
1. REST API or gRPC
2. CLI that talks with the API, preferably using [spf13/cobra](https://github.com/spf13/cobra); (as CLI framework)

### User stories/Use cases: 
1. User should be able to get questions with a number of answers

2. User should be able to select just one answer per question.

3. User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.

4. User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"


### BUILDING THE BINARY
```
go build -o quiz ./cmd
```

#### CLI commands

``` shell
quiz server # To start the api server
quiz cli # To view usage and various commands
quiz cli register # save user progress to local db
quiz cli questions # view all questions
quiz cli start # start or continues from where I left off
quiz cli answer -a A  # answer question by choosing option "A"
quiz cli stats # Result statistics
```

#### IMPROVEMENTS
- Add test cases