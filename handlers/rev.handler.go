package handlers

import (
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	svc *services.ReviewService
}

func NewReviewHandler(svc *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{svc: svc}
}

// Método para adicionar uma nova review
func (h *ReviewHandler) AddReview(c *fiber.Ctx) error {
	var body struct {
		ComicID string `json:"comic_id"`
		UserID  string `json:"user_id"`
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	review, err := h.svc.AddReview(body.UserID, body.ComicID, body.Rating, body.Comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(review)
}

// Método para buscar todas as reviews de um quadrinho específico
func (h *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	comicID := c.Params("comicID")

	reviews, err := h.svc.GetReviewsByComic(comicID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(reviews)
}

