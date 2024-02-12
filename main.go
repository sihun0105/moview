package main

import (
	"fmt"
	"server/user"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func Routers(app *fiber.App){
	app.Get("/", hello)
	app.Get("/user" ,user.GetUsers)
}
func main()  {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
	user.InitialMigration()
	app := fiber.New()
	Routers(app)
    app.Listen(":3000")
}