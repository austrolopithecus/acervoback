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

// Criar solicitação
func (h *ExchangeHandler) RequestExchange(c *fiber.Ctx) error {
	var body requests.ExchangeRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	exchange, err := h.svc.RequestExchange(body.ComicID, body.RequesterID, body.OwnerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(exchange)
}

// Aceitar solicitação
func (h *ExchangeHandler) AcceptExchange(c *fiber.Ctx) error {
	exchangeID := c.Params("id")
	userID := c.Locals("userID").(string)

	err := h.svc.AcceptExchange(exchangeID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Exchange accepted"})
}

// Recusar solicitação
func (h *ExchangeHandler) DeclineExchange(c *fiber.Ctx) error {
	exchangeID := c.Params("id")
	userID := c.Locals("userID").(string)

	err := h.svc.DeclineExchange(exchangeID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Exchange declined"})
}

func (e *ExchangeHandler) ListExchanges(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	excs, err := e.svc.ListExchanges(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(excs)
}
