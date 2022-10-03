package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("pong")
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":80",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
