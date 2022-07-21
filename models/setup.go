package models

import (
	"fmt"
	"log"

	"github.com/awesome-sphere/as-general/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUser := utils.GetenvOr("POSTGRES_USER", "pkinwza")
	dbPassword := utils.GetenvOr("POSTGRES_PASSWORD", "securepassword")
	dbHost := utils.GetenvOr("POSTGRES_HOST", "localhost")
	dbPort := utils.GetenvOr("POSTGRES_PORT", "5432")
	dbName := utils.GetenvOr("POSTGRES_DB", "as-cinema")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	var shouldSeed bool

	if !db.Migrator().HasTable("movies") {
		shouldSeed = true
	}

	db.AutoMigrate(&Movie{})
	DB = db

	if shouldSeed {
		go SeedData()
	}

	DB = db
}
