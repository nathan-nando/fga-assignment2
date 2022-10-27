package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	host, user, password, dbPort, dbName string
	db                                   *gorm.DB
	err                                  error
}

var database = Database{
	host:     "localhost",
	user:     "postgres",
	password: "pass1947",
	dbPort:   "5432",
	dbName:   "assignment2",
	db:       nil,
	err:      nil,
}

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", database.host, database.user, database.password, database.dbName, database.dbPort)

	database.db, database.err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if database.err != nil {
		log.Fatal("error connecting to database : ", database.err)
	}

	database.db.Debug().AutoMigrate(Order{}, Item{})
}

func GetDB() *gorm.DB {
	database := database.db
	return database
}
