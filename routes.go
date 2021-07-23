package main

import (
	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/book-api/controller"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/books", controller.SaveBook).Methods("POST")
	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id:[0-9]+}", controller.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{id:[0-9]+}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id:[0-9]+}", controller.DeleteBook).Methods("DELETE")
}
