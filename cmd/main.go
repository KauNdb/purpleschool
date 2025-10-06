package main

import (
	"fmt"
	"http/configs"
	"http/internal/verify"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	config := configs.LoadConfig()
	verify.NewEmail(router, config)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is started")
	server.ListenAndServe()
}
