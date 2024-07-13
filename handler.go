package quiz

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

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
		wri.WriteHeader(412)

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
		wri.WriteHeader(404)

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

	_, err = wri.Write(res)
	if err != nil {
		log.Println("err writing bytes", err)
	}
}

func answerQuestion(wri http.ResponseWriter, req *http.Request) {
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
		wri.WriteHeader(412)

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
		wri.WriteHeader(404)

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

	// Data: { result: "correct" or "wrong"}
	response := Response{
		Message: "question answered successfully",
		Data:    result,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed Conversion", err)

	}

	_, err = wri.Write(res)
	if err != nil {
		log.Println("err writing bytes", err)
	}
}
