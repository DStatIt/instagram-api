package main

import (
	"log"
	"net/http"

	"github.com/DStatIt/instagram-api/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/likes/{media_id}", handler.LikeHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
