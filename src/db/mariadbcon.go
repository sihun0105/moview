package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
func InitDb() *gorm.DB {
	DB = ConnectDatabase()
	return DB
}

func getDNS() string {
	dns := os.Getenv("DATABASE_DNS")
	fmt.Println("dns")
	fmt.Println(dns)
	if dns == "" {
		panic("Cannot get environment")
	}
	return dns
}
func ConnectDatabase() *gorm.DB{
	var err error
	
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Database connection failed: %v", r)
		}
	}()
	DB, err = gorm.Open(mysql.Open(getDNS()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	
	sqlDB, err := DB.DB()
	if err != nil {
	    log.Fatalf("Failed to get *sql.DB from *gorm.DB: %v", err)
    }

	sqlDB.SetMaxIdleConns(5)  // idle connection pool(유휴 연결 풀)의 최대 수 설정
	sqlDB.SetMaxOpenConns(20) // 데이터베이스에 대한 열결 
	sqlDB.SetConnMaxLifetime(time.Minute) // 연결의 최대 유지 시간을 1분으로 설정
	sqlDB.SetConnMaxIdleTime(time.Second * 30) // 유휴 상태의 연결이 폐기되기까지의 최대 시간을 30초로 설정

	if sqlDB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
		os.Exit(0)
	}
	return DB
}
