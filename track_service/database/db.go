package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"track_service/config"
	"track_service/models"
)

func NewConnection(conf *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)

	log.Println("Creating new database connection for track_service")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Track{})

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}
