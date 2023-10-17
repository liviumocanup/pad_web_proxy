package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"user_service/models"

	"user_service/config"
)

func NewConnection(conf *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	log.Println("Creating new database connection for user_service")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}
