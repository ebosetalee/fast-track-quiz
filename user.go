package quiz

import (
	"encoding/json"
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

