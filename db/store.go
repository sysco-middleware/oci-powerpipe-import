package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	// host will change on cloud run to /cloudsql/connection_name
	connStr       = "user=%s password=%s database=%s host=%s sslmode=disable port=%d"
	dbHost        = "host"
	dbUserKey     = "user"
	dbPassKey     = "pass"
	dbPortKey     = "port"
	dbDatabaseKey = "schema"
)

var errMissingEnvProps = errors.New("invalid.db.properties")

type Service struct {
	db *sqlx.DB
}

func New(user, pass, database, host string, port int) (*Service, string, error) {
	s := new(Service)

	conn, field, err := prepareConnString(user, pass, database, host, port)
	if err != nil {
		return nil, field, err
	}

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, "db.open", err
	}

	err = db.Ping()
	if err != nil {
		return nil, "db.ping", err
	}
	s.db = db
	return s, "", nil
}

func prepareConnString(user, pass, database, host string, port int) (string, string, error) {
	if host == "" {
		return "", dbHost, errMissingEnvProps
	}
	if user == "" {
		return "", dbUserKey, errMissingEnvProps
	}
	if pass == "" {
		return "", dbPassKey, errMissingEnvProps
	}
	if database == "" {
		return "", dbDatabaseKey, errMissingEnvProps
	}
	if port == 0 {
		return "", dbPortKey, errMissingEnvProps
	}

	return fmt.Sprintf(connStr, user, pass, database, host, port), "", nil
}
