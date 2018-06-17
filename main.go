package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	SizeB   = 1
	SizeKiB = SizeB * 1024
	SizeMiB = SizeKiB * 1024
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.Handler(http.HandlerFunc(handler)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 512 * SizeKiB,
	}

	log.Println("Serving on :8080")
	log.Fatal(s.ListenAndServe())
}
