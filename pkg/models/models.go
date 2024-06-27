package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var ErrNoRecord = errors.New("models: подходящей записи не найдено!")

// Config
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&OrderProduct{}, &Order{}, &User{}, &Product{}, &Vendor{}, &Category{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}

// func ExecuteTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
// 	transaction := db.Begin()
// 	if transaction.Error != nil {
// 		return transaction.Error
// 	}

// 	if err := fn(transaction); err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	return transaction.Commit().Error
// }
