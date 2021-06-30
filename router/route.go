package router

import (
	"github.com/gorilla/mux"
	"github.com/Akhenaten-Sama/go-gres/middleware"
	
	 
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", middleware.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", middleware.GetBook).Methods("GET")
	router.HandleFunc("/update/{id}", middleware.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/add", middleware.AddBook).Methods("POST")
	router.HandleFunc("/delete/{id}", middleware.DeleteBook).Methods("PUT")
     
	return router
}