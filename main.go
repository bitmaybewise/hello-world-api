package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/healthcheck", healthcheck)

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm healthy!")
}
