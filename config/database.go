package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnetion() (*sql.DB, error) {
	dbServer := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go"
	db, err := sql.Open(dbServer, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
