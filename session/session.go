package session

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func Wrap(handler http.Handler) http.Handler {
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * 60 * 60
	sessionManager.Cookie.Name = "session_id"
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Path = "/"
	return sessionManager.LoadAndSave(handler)
}
