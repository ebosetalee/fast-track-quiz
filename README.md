# QUIZ

This is a simple quiz about "   " with a few questions and a few alternatives for each question. Each with one correct answer. 

### Preferred Stack:
Backend - Golang
Database - Just in-memory, so no database 

### Preferred Components: 
REST API or gRPC
CLI that talks with the API, preferably using https://github.com/spf13/cobra ;( as CLI framework )

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
quiz server
quiz cli
quiz cli register # save user progress to local file
quiz cli questions
quiz cli  start # start continues from where I left off
quiz cli answer -q 1 -a A
quiz cli stats
```