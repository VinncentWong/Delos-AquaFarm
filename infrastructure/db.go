package infrastructure

import (
	"fmt"
	"os"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = _db
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func Migrate() error {
	err := db.AutoMigrate(
		&domain.Farm{},
		&domain.Pond{},
		&domain.RecordApi{},
	)
	return err
}
