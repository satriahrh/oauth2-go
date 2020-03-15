package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerInterface interface {
	Authenticate(handler http.Handler) http.Handler
	HandleAuthentication(router *mux.Router)
	PathPrefix() string
}

func FinalHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Success")
}
