package authorization_code

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	pathPrefix string
}

func NewHandler() *Handler {
	return &Handler{
		pathPrefix: "/authorization-code",
	}
}

func (h *Handler) HandleAuthentication(router *mux.Router) {

}

func (h *Handler) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) PathPrefix() string {
	return h.pathPrefix
}
