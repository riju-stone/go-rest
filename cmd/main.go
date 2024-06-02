package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/riju-stone/go-rest-api/cmd/api"
	"github.com/riju-stone/go-rest-api/config"
	"github.com/riju-stone/go-rest-api/db"
)

func connectDB(db *sql.DB) {
	error := db.Ping()

	if error != nil {
		log.Fatal(error)
	}
}

func main() {
	db, error := db.CreateNewSqlStorage(mysql.Config{
		User:                 config.EnvData.DBUser,
		Passwd:               config.EnvData.DBPassword,
		Addr:                 config.EnvData.DBAddr,
		DBName:               config.EnvData.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if error != nil {
		log.Fatal(error)
	}

	connectDB(db)

	server := api.CreateNewServer(config.EnvData.Port, nil)
	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
