package routes

import (
	"github.com/gorilla/mux"
	"github.com/cecilcodespython/bookstore-crud-api/pkg/controllers"
)

var AllRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{book_id}",controllers.GetBookbyId).Methods("GET")
	router.HandleFunc("/book/{book_id}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}",controllers.DeleteBook).Methods("DELETE")
}

