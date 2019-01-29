package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tylrd/aegis/services/user/pkg/controller"
	"github.com/tylrd/aegis/services/user/pkg/middleware"
)

type Server struct {
	Bind   string `toml:"Bind"`
	Port   int    `toml:"Port"`
	CORS   bool   `toml:"CORS"`
	Origin string `toml:"Origin"`
}

var r *mux.Router

func (s *Server) LoadRoutes() {
	r = mux.NewRouter()
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/logout", controller.Logout)
	r.HandleFunc("/session", controller.GetSession)

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.CheckAuth)
	apiRouter.HandleFunc("/user", controller.CreateUser)
}

func (s *Server) Serve() {
	var handler http.Handler

	if s.CORS {
		handler = corsHandler(r, s.Origin)
		log.Info("Starting server with CORS support")
	} else {
		handler = r
	}

	addr := fmt.Sprintf("%s:%d", s.Bind, s.Port)

	srv := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Started on port %s", addr)
	log.Debugf("Started with options: %+v", s)

	log.Fatal(srv.ListenAndServe())
}

func corsHandler(r *mux.Router, origin string) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{origin}),
		handlers.AllowCredentials(),
	)(r)
}
