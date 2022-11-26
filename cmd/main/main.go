package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/cecilcodespython/bookstore-crud-api/pkg/routes"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.AllRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8000",r))
}