package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"user_service/config"
)

func NewConnection(conf *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}
