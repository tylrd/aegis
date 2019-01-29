package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tylrd/aegis/services/user/session"

	"github.com/BurntSushi/toml"

	log "github.com/sirupsen/logrus"
	"github.com/tylrd/aegis/services/user/database"
	"github.com/tylrd/aegis/services/user/server"
)

type Config struct {
	LogLevel string                   `toml:"LogLevel"`
	Database *database.DatabaseConfig `toml:"Database"`
	Server   *server.Server           `toml:"Server"`
	Session  *session.SessionConfig   `toml:"Session"`
}

func Load(path string) *Config {
	var absPath string
	var tomlData []byte
	var err error

	if absPath, err = filepath.Abs(path); err != nil {
		log.Fatalln(err)
	}

	if tomlData, err = ioutil.ReadFile(absPath); err != nil {
		log.Fatalln(err)
	}

	var conf Config

	if _, err := toml.Decode(string(tomlData), &conf); err != nil {
		log.Fatalln(err)
	}

	lvl, ok := os.LookupEnv("LOG_LEVEL")

	if ok {
		conf.LogLevel = lvl
	}

	return &conf
}
