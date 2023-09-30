package database

import (
	"fmt"
	"log"

	"github.com/Jagadish28/go-twitter/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Check if the database exists
	var exists bool
	err = db.Raw("SELECT EXISTS (SELECT datname FROM pg_database WHERE datname = ?)", "twitter").Scan(&exists).Error
	if err != nil {
		panic("Failed to check if the database exists: " + err.Error())
	}

	// If the database doesn't exist, create it
	if !exists {
		err = db.Exec("CREATE DATABASE twitter").Error
		if err != nil {
			panic("Failed to create the database: " + err.Error())
		}
		fmt.Println("Database 'twitter' created.")
	}

	// Connect to the specific database
	dsn = "user=postgres password=admin host=localhost port=5432 dbname=twitter sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}

	// Auto-migrate the schema
	db.AutoMigrate(&model.Tweet{}, &model.Comment{})

	log.Println("Database migration completed.")
}

func Connect() {
	dsn := "host=localhost user=postgres password=admin dbname=twitter port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}
	DB = db
	log.Println("Database connection established.")

}
