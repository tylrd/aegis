package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tylrd/aegis/services/user/pkg/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	s := session.Session(r)
	uid := s.Values["user_id"]

	log.WithFields(log.Fields{
		"s_id": s.ID,
	}).Infof("User %d logged into the app", uid)

	w.WriteHeader(201)
}
