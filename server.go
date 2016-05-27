package IGapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/test/{test}", TestHandler)
	http.Handle("/", router)
	// router.HandleFunc("/instagram/test/{media_id}", handler.LikeHandler)

}
