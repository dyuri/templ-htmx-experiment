package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dyuri/templ-counter/components"
	"github.com/dyuri/templ-counter/models"
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

	// api
	mux.HandleFunc("/api/card", dh.ApiCard)

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
	card := models.Card{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)

		card.Name = r.FormValue("name")
		card.Email = r.FormValue("email")
	}

	component := components.Card(&card)
	component.Render(r.Context(), w)
}

// ApiCard serves the card as JSON
func (h *DefaultHandler) ApiCard(w http.ResponseWriter, r *http.Request) {
	card := models.Card{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)

		card.Name = r.FormValue("name")
		card.Email = r.FormValue("email")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

// Get handles GET requests
func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - GET"))
}

// Post handles POST requests
func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - POST"))
}
