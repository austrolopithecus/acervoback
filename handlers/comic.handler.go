package handlers

import (
	"acervoback/models/requests"
	"acervoback/models/responses"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type ComicHandler struct {
	svc *services.ComicService
}

func NewComicHandler(svc *services.ComicService) *ComicHandler {
	return &ComicHandler{svc: svc}
}

// Método para criar um comic
func (h *ComicHandler) CreateComic(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var body requests.NewComicRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{
			Message: "Body inválido",
			Success: false,
		})
	}

	comic, err := h.svc.Create(userID, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.CommonResponse{
			Message: "Erro ao criar comic",
			Success: false,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comic)
}

// Método para buscar comics de um usuário
func (h *ComicHandler) GetComics(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	// Recuperar os comics do serviço
	comics, err := h.svc.GetComics(userID)
	if err != nil {
		log.Err(err).Str("userID", userID).Msg("Erro ao buscar comics")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao buscar comics",
			"success": false,
		})
	}

	// Retornar os comics
	return c.Status(fiber.StatusOK).JSON(comics)
}
