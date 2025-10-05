package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type HandlerStruct struct{}

func NewHandler(router *http.ServeMux) {
	handler := &HandlerStruct{}
	router.HandleFunc("/new", handler.Response())
}

func (handler *HandlerStruct) Response() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(6) + 1
		w.Write([]byte(strconv.Itoa(num)))
	}
}
