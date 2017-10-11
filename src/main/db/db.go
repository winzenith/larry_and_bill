package db

import (
	"github.com/jackc/pgx"
)

func init() {
	connConfig := pgx.ConnConfig{
		Host:     "localhost",
		User:     "root",
		Password: "root",
		Database: "root",
	}

	_conn, err := pgx.Connect(connConfig)
	if err != nil {
		panic(err)
	}

	conn = _conn
}

var conn *pgx.Conn

// GetConn : get db instance
func GetConn() *pgx.Conn {
	return conn
}
