package main

import (
	"envelope/middleware"
	"envelope/router"
	"fmt"
	"log"
	"net/http"
)

// TODO user input validation
func main() {
	middleware.InitDb()
	fmt.Println("Db initialized")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}

