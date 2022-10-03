package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=example sslmode=disable", "db", 5432, "root", "p@ssword"))
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("err=%v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("pong")
	})
	r.HandleFunc("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		// This is BANBUTSU handler. SUPER BAD EXAMPLE

		// handling http request
		params := mux.Vars(r)
		userIDStr, ok := params["user_id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println(userID)

		// query to DB
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":80",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
