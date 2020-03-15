package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerInterface interface {
	Authenticate(handler http.Handler) http.Handler
	HandleAuthentication(router *mux.Router)
	PathPrefix() string
}
