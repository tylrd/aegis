package controller

import (
	"golang.org/x/crypto/bcypt
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"github.com/tylrd/aegis/services/user/pkg/database"
	"github.com/tylrd/aegis/services/user/pkg/model"
	"github.com/tylrd/aegis/services/user/pkg/session"
)

func Login(w http.ResponseWriter, r *http.Request) {
	s := session.Session(r)

	if r.Method == "POST" {

		var u model.User
		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}

		row := database.SQL.QueryRow("SELECT * FROM users WHERE email = ?", u.Email)
		err = row.Scan(&u)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User/Password not found", http.StatusNotFound)
			} else {
				log.Fatalln(err)
			}
		}


		if u.Password == "password" {

			res := &model.User{Email: u.Email, Username: u.Username, Role: "ADMIN"}

			js, err := json.Marshal(res)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			s.Values["user_id"] = 1
			s.Options = &sessions.Options{
				HttpOnly: true,
				Path:     "/",
				MaxAge:   86400 * 7,
			}
			s.Save(r, w)

			log.WithFields(log.Fields{
				"s_id": s.ID,
			}).Infof("User %s logged into the app", u.Username)

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}

		w.WriteHeader(403)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
