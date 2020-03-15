package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/satriahrh/oauth2-go/handler"
	"github.com/satriahrh/oauth2-go/handler/authorization-code"
	"github.com/satriahrh/oauth2-go/handler/client-credentials"
	"github.com/satriahrh/oauth2-go/handler/implicit"
	"github.com/satriahrh/oauth2-go/handler/password-credentials"
)

func final(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Success")
}

type Handler struct {
	Handler handler.HandlerInterface
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	authRouter := router.Path("/auth").Subrouter()

	for _, h := range []handler.HandlerInterface{
		implicit.NewHandler(),
		authorization_code.NewHandler(),
		password_credentials.NewHandler(),
		client_credentials.NewHandler(),
	} {
		hRouter := router.PathPrefix(h.PathPrefix()).Subrouter()
		hRouter.Use(h.Authenticate)
		hRouter.Handle("/", http.HandlerFunc(final))

		hAuthRouter := authRouter.PathPrefix(h.PathPrefix()).Subrouter()
		h.HandleAuthentication(hAuthRouter)

		readTimeout, err := time.ParseDuration(os.Getenv("READ_TIMEOUT"))
		if err != nil {
			readTimeout = 1 * time.Second
		}

		writeTimeout, err := time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
		if err != nil {
			writeTimeout = 5 * time.Second
		}

		logger := log.New(os.Stdout, "", 0)
		server := &http.Server{
			Addr:         fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT")),
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			ErrorLog:     logger,
		}
		log.Fatal(server.ListenAndServe())
	}
}
