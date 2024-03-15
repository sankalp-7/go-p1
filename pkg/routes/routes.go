package routes

import (
	"github.com/gorilla/mux"
)

var bookapprouter = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
	router.HandleFunc("/book/", controllers.getAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")
}
