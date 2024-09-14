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

	// Verifique se os campos corretos estão presentes no request
	exchange, err := h.svc.RequestExchange(body.ComicIDFrom, body.ComicIDTo, body.UserIDFrom, body.UserIDTo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(exchange)
}

func (h *ExchangeHandler) CompleteExchange(c *fiber.Ctx) error {
	exchangeID := c.Params("id")
	userID := c.Locals("userID").(string)
	err := h.svc.CompleteExchange(exchangeID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Exchange completed"})
}

