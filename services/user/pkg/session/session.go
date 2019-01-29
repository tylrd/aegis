package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	redistore "gopkg.in/boj/redistore.v1"
)

var (
	Store *redistore.RediStore
	Name  string
)

type SessionConfig struct {
	Options   sessions.Options `toml:"Options"`
	Name      string           `toml:"Name"`
	SecretKey string           `toml:"SecretKey"`
	Store     StoreConfig      `toml:"Store"`
}

type StoreConfig struct {
	Address string `toml:"Address"`
}

func Configure(s *SessionConfig) {
	Store, _ = redistore.NewRediStore(10, "tcp", s.Store.Address, "", []byte(s.SecretKey))
	Store.Options = &s.Options
	Name = s.Name
}

func Session(r *http.Request) *sessions.Session {
	s, _ := Store.Get(r, Name)
	return s
}

func IsLoggedIn(session *sessions.Session) bool {
	return session.Values["user_id"] != nil
}
