package main

import (
	"log"
	"net/http"

	"github.com/DStatIt/instagram-api/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/instagram/test/", handler.TestHandler)
	// r.HandleFunc("/instagram/follow/{account_id}", handler.FollowHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
