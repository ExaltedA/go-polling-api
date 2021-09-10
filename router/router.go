package router

import (
	"questions/handlers"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/question", handlers.GetAllQuestion).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/question", handlers.CreateQuestion).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteItem", handlers.DeleteQuestion).Methods("DELETE", "OPTIONS")
	
	return router
}
