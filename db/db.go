package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func CreateNewSqlStorage(config mysql.Config) (*sql.DB, error) {
	db, error := sql.Open("mysql", config.FormatDSN())
	if error != nil {
		log.Fatal(error)
	}

	return db, nil
}
