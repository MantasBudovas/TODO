package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"app/internal/entity"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Card{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
