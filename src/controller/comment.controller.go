package controller

import (
	"moview/src/models"
	"moview/src/repository"

	"github.com/gofiber/fiber/v2"
)

type CommentController struct {
	CommentRepository repository.CommentRepository
}

func NewCommentController(commentRepository repository.CommentRepository) *CommentController {
	return &CommentController{CommentRepository: commentRepository}
}

func (cc *CommentController) CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return err
	}

	if err := cc.CommentRepository.CreateComment(comment); err != nil {
		return err
	}
	return c.JSON(comment)
}

func (cc *CommentController) GetCommentByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	comment, err := cc.CommentRepository.GetCommentByID(id)
	if err != nil {
		return err
	}
	return c.JSON(comment)
}

func (cc *CommentController) UpdateComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return err
	}
	comment.ID = id

	if err := cc.CommentRepository.UpdateComment(comment); err != nil {
		return err
	}
	return c.JSON(comment)
}

func (cc *CommentController) DeleteComment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	if err := cc.CommentRepository.DeleteComment(id); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}