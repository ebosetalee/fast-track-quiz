package quiz

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getAllQuestions(wri http.ResponseWriter, req *http.Request) {
	wri.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		wri.WriteHeader(http.StatusMethodNotAllowed)

		response := Response{
			Error: "Whoops! Route Method Wrong.",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)

		}

		_, err = wri.Write(res)
		if err != nil {
			log.Println("failed to writes response header", err)
		}

		return
	}

	// sort quiz questions in place
	Quiz.sort()

	response := Response{
		Message: "questions retrieved successfully",
		Data:    Quiz,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed Conversion", err)

	}

	wri.WriteHeader(http.StatusOK)
	_, err = wri.Write(res)
	if err != nil {
		log.Println("err writing bytes", err)
	}
}

func getQuestion(wri http.ResponseWriter, req *http.Request) {
	wri.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodGet {
		wri.WriteHeader(http.StatusMethodNotAllowed)

		response := Response{
			Error: "Whoops! Route Method Wrong.",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)

		}

		_, err = wri.Write(res)
		if err != nil {
			log.Println("failed to writes response header", err)
		}

		return
	}

	index := req.PathValue("id")

	number, err := strconv.Atoi(index)
	if err != nil {
		wri.WriteHeader(http.StatusPreconditionFailed)

		response := Response{
			Error: "Id is not a valid number",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)
		}

		_, err = wri.Write(res)
		if err != nil {
			log.Println("failed to writes response header", err)
		}

		return
	}

	q := Quiz.getQuestion(number)
	if q == nil {
		wri.WriteHeader(http.StatusNotFound)

		// return an error
		response := Response{
			Error: "Question does not exist",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)
		}

		_, err = wri.Write(res)
		if err != nil {
			log.Println("failed to writes response header", err)
		}
		return
	}

	response := Response{
		Message: "question retrieved successfully",
		Data:    q,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed Conversion", err)

	}

	wri.WriteHeader(http.StatusOK)
	_, err = wri.Write(res)
	if err != nil {
		log.Println("err writing bytes", err)
	}
}

func answerQuestion(db *Database) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
		wri.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodPost {
			wri.WriteHeader(http.StatusMethodNotAllowed)

			response := Response{
				Error: "Whoops! Route Method Wrong.",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)

			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}

			return
		}

		index := req.PathValue("id")

		number, err := strconv.Atoi(index)
		if err != nil {
			wri.WriteHeader(http.StatusPreconditionFailed)

			response := Response{
				Error: "Id is not a valid number",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)
			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}

			return
		}

		q := Quiz.getQuestion(number)
		if q == nil {
			wri.WriteHeader(http.StatusNotFound)

			// return an error
			response := Response{
				Error: "Question does not exist",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)
			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}
			return
		}

		var answer UserAnswer
		if err := ReadJSON(req, &answer); err != nil {
			log.Println("Failed to parse request body", err)

			wri.WriteHeader(http.StatusBadRequest)

			// return an error
			response := Response{
				Error: "Failed to parse request body",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)
			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}
			return
		}

		result := q.checkAnswer(answer.ChosenAnswer)


		// use db to update the users score
		// db.addScore() ???
		getUserID, ok := req.Header[http.CanonicalHeaderKey("x-quiz-userId")]
		if !ok {
			wri.WriteHeader(http.StatusUnauthorized)

			// return an error
			response := Response{
				Error: "unauthorized",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)
			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}
			return
		}

		// check user exists
		userId := getUserID[0]
		_, ok = db.Users[userId]
		if !ok {
			wri.WriteHeader(http.StatusNotFound)

			// return an error
			response := Response{
				Error: "User does not exist, kindly register to answer",
			}

			res, err := json.Marshal(response)
			if err != nil {
				log.Println("Failed Conversion", err)
			}

			_, err = wri.Write(res)
			if err != nil {
				log.Println("failed to writes response header", err)
			}
			return
		}


		if result == "Correct"{
			db.addScore(userId)
		}

		// Data: { result: "Correct" or "Wrong"}
		response := Response{
			Message: "question answered successfully",
			Data:    result,
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)

		}

		wri.WriteHeader(http.StatusOK)
		_, err = wri.Write(res)
		if err != nil {
			log.Println("err writing bytes", err)
		}
	}
}
