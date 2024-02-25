package controller

import (
	"moview/src/repository"

	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	MovieRepository repository.MovieRepository
}

func NewMovieController(movieRepository repository.MovieRepository) *MovieController {
	return &MovieController{MovieRepository: movieRepository}
}

func (mc *MovieController) GetMovieByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	movie, err := mc.MovieRepository.GetMoviewByID(id)
	if err != nil {
		return err
	}
	return c.JSON(movie)
}


func (mc *MovieController) GetMovies(c *fiber.Ctx) error {
	movies, err := mc.MovieRepository.GetMovies()
	if err != nil {
		return err
	}
	return c.JSON(movies)
}

func (mc *MovieController) FetchMovies(c *fiber.Ctx) error {
	date := c.Params("date")
	if err := mc.MovieRepository.FetchMovies(date); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}