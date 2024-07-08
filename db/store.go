package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

const (
	// host will change on cloud run to /cloudsql/connection_name
	connStr       = "user=%s password=%s database=%s host=%s sslmode=disable port=%s"
	dbHost        = "PG_DB_HOST"
	dbUserKey     = "PG_DB_USER"
	dbPassKey     = "PG_DB_PASS"
	dbPortKey     = "PG_DB_PORT"
	dbDatabaseKey = "PG_DB_SCHEMA"
)

var errMissingEnvProps = errors.New("missing environment variables")

type Service struct {
	db *sqlx.DB
}

func New() (*Service, string, error) {
	s := new(Service)

	conn, field, err := prepareConnString()
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

func prepareConnString() (string, string, error) {
	host := os.Getenv(dbHost)
	if host == "" {
		return "", dbHost, errMissingEnvProps
	}

	user := os.Getenv(dbUserKey)
	if user == "" {
		return "", dbUserKey, errMissingEnvProps
	}

	pass := os.Getenv(dbPassKey)
	if pass == "" {
		return "", dbPassKey, errMissingEnvProps
	}

	database := os.Getenv(dbDatabaseKey)
	if database == "" {
		return "", dbDatabaseKey, errMissingEnvProps
	}

	port := os.Getenv(dbPortKey)
	if port == "" {
		return "", dbPortKey, errMissingEnvProps
	}

	return fmt.Sprintf(connStr, user, pass, database, host, port), "", nil
}
