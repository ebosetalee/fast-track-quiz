package quiz

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database := NewDatabase()

	// define handler
	var mux = http.NewServeMux()

	// http methods
	mux.HandleFunc("/quiz/questions", getAllQuestions)

	mux.HandleFunc("/quiz/{id}", getQuestion)

	mux.HandleFunc("/quiz/{id}/answer", answerQuestion(database))

	mux.HandleFunc("/quiz/user/register", registerUser(database))

	mux.HandleFunc("/quiz/user/stats", checkStats(database))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server started on port %s\n", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
