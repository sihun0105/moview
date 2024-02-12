package user

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const envDNS = "DATABASE_DNS"

func getDNS() string {
	dns := os.Getenv(envDNS)
	fmt.Println("dns")
	fmt.Println(dns)
	if dns == "" {
		panic("Cannot get environment")
	}
	return dns
}

type User struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(getDNS()), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error{
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}