package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"ccproject/service"
	"ccproject/domain"
	"github.com/gorilla/mux"
	"strconv"
	
)

type BookHandler struct {
	service *service.BooksService
}

func NewBookHandler(service *service.BooksService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}


func (handler *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := handler.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResponse(books, w)
}

func (handler *BookHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rez, err := handler.service.Borrow(&book)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rez))
}


func (handler *BookHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rez, err := handler.service.Register(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rez))
}

func (handler *BookHandler) Return(w http.ResponseWriter, r *http.Request) {
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

	isbn, ok := vars["isbn"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
	rez, err := handler.service.Return(userID, isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rez))
}



