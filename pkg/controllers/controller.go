package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/cecilcodespython/bookstore-crud-api/pkg/models"
	"github.com/cecilcodespython/bookstore-crud-api/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	response ,_ := json.Marshal(newBooks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookbyId(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	bookID := parameters["book_id"]
	bookIDNum,err := strconv.ParseInt(bookID,0,0)
	if err != nil {
		log.Fatal(err)
	}
	bookDetails,_ := models.GetBookbyId(bookIDNum)
	response,_ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.Parser(r,CreateBook)
	b := CreateBook.CreateBook()
	response,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	bookID := parameters["book_id"]
	bookIdNum,err:= strconv.ParseInt(bookID,0,0)
	if err != nil {
		log.Fatal(err)
	}
	book := models.DeleteBook(bookIdNum)
	response,_ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook = &models.Book{}
	utils.Parser(r,UpdateBook)
	parameters := mux.Vars(r)
	bookId := parameters["book_id"]
	bookIdNum,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		log.Fatal(err)
	}
	bookDetails,db := models.GetBookbyId(bookIdNum)
	if UpdateBook.Title != ""{
		bookDetails.Title = UpdateBook.Title
	}
	if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publisher != ""{
		bookDetails.Publisher = UpdateBook.Publisher
	}
	if UpdateBook.Isbn != ""{
		bookDetails.Isbn = UpdateBook.Isbn
	}
	db.Save(&bookDetails)
	response ,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}