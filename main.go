package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MadAppGang/httplog"
	"github.com/dyuri/templ-counter/db"
	"github.com/dyuri/templ-counter/handlers"
	"github.com/dyuri/templ-counter/services"
	"github.com/dyuri/templ-counter/session"
	"golang.org/x/exp/slog"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	store, err := db.NewCountStore()
	if err != nil {
		log.Error("failed to create store", err)
		os.Exit(1)
	}

	counterService := services.NewCounter(log, store)
	handler := handlers.NewHandler(log, counterService)

	handlerWithSession := session.Wrap(handler)

	server := &http.Server{
		Addr:         ":8000",
		Handler:      httplog.Logger(handlerWithSession),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Server is listening on %v\n", server.Addr)
	server.ListenAndServe()
}
