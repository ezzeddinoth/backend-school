package database

import (
	"fmt"
	"log"
	"os"
	"school/internal/database/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Repository *repository.Repository

func Connect() (err error) {

	Db, err = gorm.Open(postgres.Open(buildConnectionString()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening GORM connection: %v", err)
	}

	log.Printf("connected to database server: %s",
		os.Getenv("SCHOOL_DB_SERVER"))
	Repository = repository.SetupRepositories(Db)
	log.Printf("database repositories initialized")

	return
}

func buildConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("SCHOOL_DB_USER"),
		os.Getenv("SCHOOL_DB_PASSWORD"),
		os.Getenv("SCHOOL_DB_SERVER"),
		os.Getenv("SCHOOL_DB_PORT"),
		os.Getenv("SCHOOL_DB_DATABASE"),
	)
}
