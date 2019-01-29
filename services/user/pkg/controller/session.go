package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tylrd/aegis/services/user/pkg/model"
	"github.com/tylrd/aegis/services/user/pkg/session"
)

func GetSession(w http.ResponseWriter, r *http.Request) {
	s := session.Session(r)

	if s == nil {
		http.Error(w, "No session found", 404)
		return
	}

	uid := s.Values["user_id"]

	if uid == nil {
		http.Error(w, "No session found", 404)
		return
	}

	u := &model.User{Username: "taylor", Role: "ADMIN"}
	res, _ := json.Marshal(u)

	w.Write([]byte(res))
}
