package main

import (
	"backend_example/db"
	"backend_example/handler"
	"backend_example/model"
	"backend_example/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.Open()
	if err != nil {
		log.Fatalf("error in connecting db%v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("pong")
	})
	r.HandleFunc("/bad/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		// This is BANBUTSU handler. SUPER BAD EXAMPLE

		// handling http request
		params := mux.Vars(r)
		userIDStr, ok := params["user_id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// query to DB
		user := model.User{}
		if err := db.Get(&user, "SELECT * FROM users WHERE id = $1", userID); err != nil {
			// HTTPの知識がここに紛れ込んでしまう
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// handling http response
		w.WriteHeader(http.StatusOK)
		body, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(body)

	})

	// 私が書くならこうするという例
	userHandler := handler.NewUser(repository.NewUser(db))
	r.HandleFunc("/users/{user_id}", userHandler.Get)

	srv := &http.Server{
		Handler: handlers.CombinedLoggingHandler(os.Stdout, r),
		Addr:    ":80",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
