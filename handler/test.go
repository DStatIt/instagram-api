package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "found", vars["media_id"])

}
