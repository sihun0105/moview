package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() (*gorm.DB, error) {
    db, err := ConnectDatabase()
    if err != nil {
        return nil, err
    }
    return db, nil
}

func ConnectDatabase() (*gorm.DB, error) {
    dbcon := getDSN()
    DB, err := gorm.Open(mysql.Open(dbcon), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %w", err)
    }
    log.Println("Connected to the database")
    return DB, nil
}

func getDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"))
	return dsn
}