package user

import (
	"moview/src/models"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const envDNS = "DATABASE_DNS"

func getDNS() string {
	dns := os.Getenv(envDNS)
	if dns == "" {
		panic("Cannot get environment")
	}
	return dns
}

func InitialMigration() {
    DB, err = gorm.Open(mysql.Open(getDNS()), &gorm.Config{})
    if err != nil {
        panic("Cannot connect DB")
    }
}

func GetUsers(c *fiber.Ctx) error{
	var users []models.User
	DB.Find(&users)
	return c.JSON(&users)
}