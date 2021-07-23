package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nitinigi2/practice/book-api/model"
	"github.com/nitinigi2/practice/book-api/repository"
)

func SaveBook(w http.ResponseWriter, r *http.Request) {

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = repository.SaveBook(book)

	if err != nil {
		log.Fatal("Error in saving book", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Book saved successfully.")
	w.WriteHeader(http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, bookId int) {

	err := repository.DeleteBook(bookId)

	if err != nil {
		log.Fatal("Error in deleting book", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Book deleted successfully.")
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, bookId int) {

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatal("Error in parsing book payload", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repository.UpdateBook(book, bookId)
	if err != nil {
		log.Fatal("Error in updating book", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Book updated successfully.")
	w.WriteHeader(http.StatusOK)
}

func GetBook(w http.ResponseWriter, bookId int) {

	w.Header().Set("Content-Type", "application/json")

	book, err := repository.GetBook(bookId)

	if err != nil {
		log.Println("Book not found with id", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := repository.GetAllBooks()
	if err != nil {
		fmt.Println("No book found", err)
	}
	json.NewEncoder(w).Encode(books)
}
