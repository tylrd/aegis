package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/tylrd/aegis/services/user/pkg/session"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	s := session.Session(r)

	if session.IsLoggedIn(s) {
		s.Options.MaxAge = -1
		if err := sessions.Save(r, w); err != nil {
			http.Error(w, "Error logging out.", 500)
			return
		}
		w.Write([]byte("Success"))
		return
	}
	http.Error(w, "No session found", 404)
}
