package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
	log "github.com/sirupsen/logrus"
)

const (
	defaultUserName = "postgres"
)

var SQL *sql.DB

type DatabaseConfig struct {
	Username  string `toml:"Username"`
	Password  string `toml:"Password"`
	Name      string `toml:"Name"`
	Hostname  string `toml:"Hostname"`
	Port      int    `toml:"Port"`
	Parameter string `toml:Parameter`
}

func Connect(c *DatabaseConfig) {

	if c.Username == "" {
		c.Username = "postgres"
	}

	if c.Name == "" {
		c.Name = "aegis"
	}

	if c.Hostname == "" {
		c.Name = "localhost"
	}

	if c.Port == 0 {
		c.Port = 5432
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s", c.Username, c.Password, c.Hostname, c.Port, c.Name, c.Parameter)

	var err error

	log.Infof("Connecting to database %s", connStr)

	if SQL, err = sql.Open("postgres", connStr); err != nil {
		log.Fatalln("SQL Driver Error", err)
	}

}
