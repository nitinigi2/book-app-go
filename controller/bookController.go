package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/book-api/service"
)

func SaveBook(w http.ResponseWriter, r *http.Request) {
	ok, role, err := service.Authorize(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if role == "admin" && ok {
		service.SaveBook(w, r)
		log.Println("Book saved successfully")
	} else {
		log.Println("User doesn't have permission to save book")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service.GetBooks(w, r)
	w.WriteHeader(http.StatusOK)
	log.Println("Books retrived successfully")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := getBooksId(r)
	if err != nil {
		log.Fatal("Id not valid")
		return
	}

	ok, role, err := service.Authorize(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if role == "admin" && ok {
		service.UpdateBook(w, r, bookId)
	} else {
		log.Fatal("User doesn't have permission to perform this action")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := getBooksId(r)
	if err != nil {
		log.Fatal("Id not valid")
		return
	}

	ok, role, err := service.Authorize(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if role == "admin" && ok {
		service.DeleteBook(w, bookId)
	} else {
		log.Fatal("User doesn't have permission to perform this action")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId, err := getBooksId(r)
	if err != nil {
		log.Fatal("Id not valid")
		return
	}
	service.GetBook(w, bookId)
}

func getBooksId(r *http.Request) (int, error) {
	params := mux.Vars(r)
	id := params["id"]
	return strconv.Atoi(id)
}
