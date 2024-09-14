package handlers

import (
	"acervoback/models/requests"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
)

type ExchangeHandler struct {
	svc *services.ExchangeService
}

func NewExchangeHandler(svc *services.ExchangeService) *ExchangeHandler {
	return &ExchangeHandler{svc: svc}
}

func (h *ExchangeHandler) RequestExchange(c *fiber.Ctx) error {
	var body requests.ExchangeRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	exchange, err := h.svc.RequestExchange(body.User1ID, body.User2ID, body.Comic1ID, body.Comic2ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(exchange)
}

func (h *ExchangeHandler) CompleteExchange(c *fiber.Ctx) error {
	exchangeID := c.Params("id")
	err := h.svc.CompleteExchange(exchangeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Exchange completed"})
}

