package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connection Failed...\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the Database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")

	//TODO: Handle Migrations
	Database = DbInstance{Db: db}
}
