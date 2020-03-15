package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerInterface interface {
	Authenticate(handler http.Handler) http.Handler
	HandleAuthentication(router *mux.Router)
	PathPrefix() string
}
