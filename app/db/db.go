package db

import (
	"fmt"
	"log"

	"battle-of-monsters/app/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	if DB != nil {
		return DB
	}

	log.Println("Database connection started")

	DB, err := gorm.Open(sqlite.Open(getDataBaseName()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to open the database connection. %w", err))
	}

	return DB
}

func getDataBaseName() string {
	dn := config.ENV.DBName
	if dn == "" {
		log.Fatalln("database name is not defined to open a connection")
	}

	return dn
}
