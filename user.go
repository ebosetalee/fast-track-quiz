package quiz

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func registerUser(db *Database) http.HandlerFunc {
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

		var registerUser UserRequest
		if err := ReadJSON(req, &registerUser); err != nil {
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

		user := User{
			Id:    registerUser.Id,
			Score: 0,
		}

		// improvement: check if user has registered, do we throw an error?
		_, ok := db.Users[user.Id]
		if ok {
			wri.WriteHeader(http.StatusConflict)

			// return an error
			response := Response{
				Error: "User already exists",
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

		db.register(user)

		response := Response{
			Message: "user created successfully",
			Data:    user,
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Println("Failed Conversion", err)

		}

		wri.WriteHeader(http.StatusCreated)
		_, err = wri.Write(res)
		if err != nil {
			log.Println("err writing bytes", err)
		}
	}
}

func checkStats(db *Database) http.HandlerFunc {
	return func(wri http.ResponseWriter, req *http.Request) {
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

		userIDHeader, ok := req.Header[http.CanonicalHeaderKey("x-quiz-userId")]
		if !ok {
			wri.WriteHeader(http.StatusNotFound)

			// return an error
			response := Response{
				Error: "No Auth provided",
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

		userId := userIDHeader[0]

		// check user exists
		_, ok = db.Users[userId]
		if !ok {
			wri.WriteHeader(http.StatusNotFound)

			// return an error
			response := Response{
				Error: "User does not exist",
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

		// get user position
		users := db.sortByScore()
		var position int
		var score int
		for i := 0; i < len(users); i++ {
			if users[i].Id == userId {
				score = users[i].Score
				position = i
				break
			}
		}

		// calculate user percentile
		percentile := float32(len(users)-position) / float32(len(users)) * 100

		fmt.Printf("User is in %.0f percentile with %v users, position %v and score %v\n", percentile, len(users), position, score)

		message := fmt.Sprintf("You were better than %.0f%% of all quizzers", percentile)

		response := Response{
			Message: "result retrieved successfully",
			Data:    message,
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
