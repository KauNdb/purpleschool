package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is started")
	server.ListenAndServe()
}
