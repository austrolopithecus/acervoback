package handlers

import (
	"acervoback/models/requests"
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
	var body requests.ReviewRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Validações básicas
	if body.Rating < 1 || body.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Rating must be between 1 and 5"})
	}

	if body.Comment == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Comment cannot be empty"})
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

