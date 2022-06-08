package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Compound(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(
		fmt.Sprintf("Hello %+v", vars),
	))
}
