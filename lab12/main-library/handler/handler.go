package handler

import (
	"ccmainproject/domain"
	"ccmainproject/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *service.UsersService
}

func NewUserHandler(service *service.UsersService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResponse(users, w)
}

func (handler *UserHandler) Borrow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}

	users, err := handler.service.Borrow(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(users, w)
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := handler.service.Register(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	responseBody := fmt.Sprintf("User ID isss: %d", userId)
	w.Write([]byte(responseBody))
}

func (handler *UserHandler) Return(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["userId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}

	users, err := handler.service.Return(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(users, w)
}
