package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/logger"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/satriahrh/oauth2-go/handler"
	"github.com/satriahrh/oauth2-go/handler/authorization-code"
	"github.com/satriahrh/oauth2-go/handler/client-credentials"
	"github.com/satriahrh/oauth2-go/handler/implicit"
	"github.com/satriahrh/oauth2-go/handler/password-credentials"
)

func main() {
	loggerIo := os.Stdout
	defer logger.Init("OAuth2 Go Logger", true, true, loggerIo).Close()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
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
		hRouter.HandleFunc("/", handler.FinalHandler)

		hAuthRouter := authRouter.PathPrefix(h.PathPrefix()).Subrouter()
		h.HandleAuthentication(hAuthRouter)
	}

	readTimeout, err := time.ParseDuration(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		readTimeout = 1 * time.Second
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		writeTimeout = 5 * time.Second
	}

	loggedRouter := handlers.LoggingHandler(loggerIo, router)
	server := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler:      loggedRouter,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	logger.Fatal(server.ListenAndServe())
}
