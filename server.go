package main

import (
	"net/http"

	"github.com/DStatIt/instagram-api/handler"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/test/{testId}", handler.TestHandler)
	// router.HandleFunc("/instagram/test/{media_id}", handler.LikeHandler)
	http.Handle("/", router)
}

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/instagram/test/{media_id}", handler.LikeHandler)
// 	// r.HandleFunc("/instagram/follow/{account_id}", handler.FollowHandler)
// 	http.Handle("/", r)
//
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }
