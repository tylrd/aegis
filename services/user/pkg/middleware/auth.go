package middleware

import (
	"net/http"

	"github.com/tylrd/aegis/services/user/pkg/session"
)

func CheckAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := session.Session(r)

		if session.IsLoggedIn(s) {
			h.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
