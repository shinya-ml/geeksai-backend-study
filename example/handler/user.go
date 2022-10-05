package handler

import (
	"backend_example/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	repository *repository.User
}

func NewUser(repo *repository.User) *User {
	return &User{
		repository: repo,
	}
}

func (h *User) Get(w http.ResponseWriter, r *http.Request) {
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
	user, err := h.repository.FindByID(userID)
	if err != nil {
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
}
