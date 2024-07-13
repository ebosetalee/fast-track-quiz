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

	// define handler
	var mux = http.NewServeMux()

	// http methods
	mux.HandleFunc("/questions", getAllQuestions)
	mux.HandleFunc("/quiz/{id}", getQuestion)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server started on port %s\n", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

