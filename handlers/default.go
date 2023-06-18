package handlers

import (
	"fmt"
	"net/http"

	"github.com/dyuri/templ-counter/components"
	"github.com/dyuri/templ-counter/services"
	"golang.org/x/exp/slog"
)

// NewHandler creates a new DefaultHandler
func NewHandler(log *slog.Logger, counter *services.Counter) http.Handler {
	mux := http.NewServeMux()

	dh := &DefaultHandler{
		Logger:  log,
		Counter: counter,
		mux:     mux,
	}

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/assets/", fs)
	mux.HandleFunc("/", dh.Index)
	mux.HandleFunc("/about", dh.About)

	// widgets
	mux.HandleFunc("/widget/card", dh.Card)

	return dh
}

// DefaultHandler is the default request handler
type DefaultHandler struct {
	Logger  *slog.Logger
	Counter *services.Counter
	mux     *http.ServeMux
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// Index serves the index page
func (h *DefaultHandler) Index(w http.ResponseWriter, r *http.Request) {
	component := components.Index()
	component.Render(r.Context(), w)
}

// About serves the about page
func (h *DefaultHandler) About(w http.ResponseWriter, r *http.Request) {
	component := components.About()
	component.Render(r.Context(), w)
}

// Card serves the card widget
func (h *DefaultHandler) Card(w http.ResponseWriter, r *http.Request) {
	name := "cica"
	email := "cica@kutya.hu"

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)

		h.Logger.Info(fmt.Sprintf("POST"))
		for k, v := range r.Form {
			h.Logger.Warn(fmt.Sprintf("key: %s, value: %s\n", k, v))
		}
		name = r.FormValue("name")
		email = r.FormValue("email")
	}

	component := components.Card(name, email)
	component.Render(r.Context(), w)
}

// Get handles GET requests
func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - GET"))
}

// Post handles POST requests
func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - POST"))
}
