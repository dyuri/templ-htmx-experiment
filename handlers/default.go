package handlers

import (
	"net/http"

	"github.com/dyuri/templ-counter/components"
	"github.com/dyuri/templ-counter/services"
	"golang.org/x/exp/slog"
)

// New creates a new DefaultHandler
func New(log *slog.Logger, counter *services.Counter) http.Handler {
	mux := http.NewServeMux()

	dh := &DefaultHandler{
		Logger:  log,
		Counter: counter,
		mux:     mux,
	}

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/assets/", fs)
	mux.HandleFunc("/", dh.Get)
	mux.HandleFunc("/get2", dh.Get2)

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

// Get handles GET requests
func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	component := components.Index()
	component.Render(r.Context(), w)
}

// Get2 handles GET requests
func (h *DefaultHandler) Get2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - GET2"))
}

// Post handles POST requests
func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World - POST"))
}
