package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	log.Println("Listening...")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        fs,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
