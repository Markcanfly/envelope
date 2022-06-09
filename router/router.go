package router

import (
	"envelope/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/message", middleware.GetAllOpenedMessages).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/message", middleware.CreateMessage).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/message/{id}", middleware.DeleteMessage).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/user/register", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/login", middleware.Login).Methods("POST", "OPTIONS")
	return router
}
