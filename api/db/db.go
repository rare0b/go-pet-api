package db

import "database/sql"

type DB interface {
}

func NewDB() *sql.DB {
	return nil
}
