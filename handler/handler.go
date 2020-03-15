package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerInterface interface {
	Authenticate(handler http.Handler) http.Handler
	HandleAuthentication(router *mux.Router)
	PathPrefix() string
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	response := struct{ Message string `json:"message"` }{
		Message: "OK",
	}
	_ = json.NewEncoder(w).Encode(response)
}
