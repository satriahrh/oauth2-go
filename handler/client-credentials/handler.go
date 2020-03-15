package client_credentials

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	pathPrefix string
}

func NewHandler() *Handler {
	return &Handler{
		pathPrefix: "/client-credentials",
	}
}

func (h *Handler) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) HandleAuthentication(router *mux.Router) {

}

func (h *Handler) PathPrefix() string {
	return h.pathPrefix
}
