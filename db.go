package utils

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {

	dbUrl := GetWithDefault("AWSTIO_DB_URL", "postgres://postgres:postgres@localhost:5432/awstdev_db?sslmode=disable")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	maxOpenConnect, err := strconv.Atoi(GetWithDefault("AWSTIO_MAX_OPEN_CONNS", "10"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnect)

	return db, nil
}
