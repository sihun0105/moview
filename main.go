package main

import (
	"fmt"
	"log"
	"moview/src/controller"
	"moview/src/db"
	"moview/src/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)
var DB *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB, err = db.InitDB()
    if err != nil {
        panic("Cannot connect DB")
    }
	app := fiber.New()

	userRepository := repository.NewUserRepository(DB)
	userController := controller.NewUserController(userRepository)

	commentRepository := repository.NewCommentRepository(DB)
	commentController := controller.NewCommentController(commentRepository)

	setupRoutes(app, userController, commentController)
	port := ":3000"
	fmt.Printf("Server is listening on port %s\n", port)
	log.Fatal(app.Listen(port))
}

func setupRoutes(
	app *fiber.App, 
	uc *controller.UserController,
	cc *controller.CommentController,
	) {
	app.Post("/users", uc.CreateUser)
	app.Get("/users/:id", uc.GetUserByID)
	app.Put("/users/:id", uc.UpdateUser)
	app.Delete("/users/:id", uc.DeleteUser)

	app.Post("/comments", cc.CreateComment)
	app.Get("/comments/:id", cc.GetCommentByID)
	app.Put("/comments/:id", cc.UpdateComment)
	app.Delete("/comments/:id", cc.DeleteComment)
	
}

