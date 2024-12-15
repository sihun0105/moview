package main

import (
	"fmt"
	"log"
	"moview/src/controller"
	"moview/src/db"
	"moview/src/lib/common"
	"moview/src/repository"
	"net/http/pprof"
	"os"
	"time"

	"github.com/gofiber/adaptor/v2"
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

	registerPprof(app)

	authRepository := repository.NewAuthRepository(DB)
	authController := controller.NewAuthController(authRepository)

	userRepository := repository.NewUserRepository(DB)
	userController := controller.NewUserController(userRepository)

	commentRepository := repository.NewCommentRepository(DB)
	commentController := controller.NewCommentController(commentRepository)

	movieRepository := repository.NewMovieRepository(DB)
	movieController := controller.NewMovieController(movieRepository)

	common.StartUpdater(func() {
		UpdateMovies(movieRepository)
	})

	setupRoutes(app, userController, commentController, movieController, authController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}
	port = ":" + port
	fmt.Printf("Server is listening on port %s\n", port)
	log.Fatal(app.Listen(port))
}

func registerPprof(app *fiber.App) {
	pprofGroup := app.Group("/debug/pprof")
	pprofGroup.Get("/", adaptor.HTTPHandlerFunc(pprof.Index))
	pprofGroup.Get("/cmdline", adaptor.HTTPHandlerFunc(pprof.Cmdline))
	pprofGroup.Get("/profile", adaptor.HTTPHandlerFunc(pprof.Profile))
	pprofGroup.Get("/symbol", adaptor.HTTPHandlerFunc(pprof.Symbol))
	pprofGroup.Get("/trace", adaptor.HTTPHandlerFunc(pprof.Trace))
	pprofGroup.Get("/heap", adaptor.HTTPHandlerFunc(pprof.Handler("heap").ServeHTTP))
	pprofGroup.Get("/goroutine", adaptor.HTTPHandlerFunc(pprof.Handler("goroutine").ServeHTTP))
	pprofGroup.Get("/threadcreate", adaptor.HTTPHandlerFunc(pprof.Handler("threadcreate").ServeHTTP))
	pprofGroup.Get("/block", adaptor.HTTPHandlerFunc(pprof.Handler("block").ServeHTTP))
}

func setupRoutes(
	app *fiber.App,
	uc *controller.UserController,
	cc *controller.CommentController,
	mc *controller.MovieController,
	ac *controller.AuthController,
) {
	app.Post("/users", uc.CreateUser)
	app.Get("/users/:id", uc.GetUserByID)
	app.Put("/users/:id", uc.UpdateUser)
	app.Delete("/users/:id", uc.DeleteUser)

	app.Post("/comments", cc.CreateComment)
	app.Get("/comments/:id", cc.GetCommentByID)
	app.Put("/comments/:id", cc.UpdateComment)
	app.Delete("/comments/:id", cc.DeleteComment)

	app.Get("/movies/:id", mc.GetMovieByID)
	app.Get("/movies", mc.GetMovies)
	app.Get("/movies/fetch/:date", mc.FetchMovies)
}

func UpdateMovies(repo repository.MovieRepository) {
	currentTime := time.Now().AddDate(0, 0, -1)

	date := currentTime.Format("20060102")

	err := repo.FetchMovies(date)
	if err != nil {
		fmt.Println("Failed to update movies:", err)
		return
	}

	fmt.Println("Movies updated successfully.")
}
