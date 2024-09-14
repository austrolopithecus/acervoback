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

// Handler para adicionar uma nova review
func (h *ReviewHandler) AddReview(c *fiber.Ctx) error {
	var body requests.ReviewRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	review, err := h.svc.AddReview(body.UserID, body.ComicID, body.Rating, body.Comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(review)
}

// Handler para obter reviews de um quadrinho
func (h *ReviewHandler) GetReviewsByComic(c *fiber.Ctx) error {
	comicID := c.Params("comicID")
	reviews, err := h.svc.GetReviewsByComic(comicID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(reviews)
}

