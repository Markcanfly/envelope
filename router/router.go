package router

import (
	"envelope/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/message", middleware.GetAllMessages).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/message", middleware.CreateMessage).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/message/{id}", middleware.DeleteMessage).Methods("DELETE", "OPTIONS")
	return router
}
