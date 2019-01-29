package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/tylrd/aegis/services/user/pkg/config"
	"github.com/tylrd/aegis/services/user/pkg/database"
	"github.com/tylrd/aegis/services/user/pkg/session"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "config.toml", "Path to configuration file. Defaults to config.toml")
	flag.Parse()

	c := config.Load(confPath)
	ll, err := log.ParseLevel(c.LogLevel)

	if err != nil {
		ll = log.InfoLevel
	}

	log.SetLevel(ll)

	database.Connect(c.Database)
	session.Configure(c.Session)

	c.Server.LoadRoutes()
	c.Server.Serve()
}
