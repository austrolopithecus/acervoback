package handlers

import (
	"acervoback/models/requests"
	"acervoback/models/responses"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
)

type ComicHandler struct {
	svc *services.ComicService
}

func NewComicHandler(svc *services.ComicService) *ComicHandler {
	return &ComicHandler{svc: svc}
}

// CreateComic godoc
// @Summary Cria um quadrinho com base no ISBN
// @Description Cria um quadrinho com base no ISBN
// @Tags Comic
// @Produce json
// @Security TokenAuth
// @Param body body requests.NewComicRequest true "Body"
// @Success 200 {object} models.Comic
// @Failure 500 {object} responses.CommonResponse
// @Failure 401 {object} responses.CommonResponse
// @Router /comic [put]
func (h *ComicHandler) CreateComic(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	var body requests.NewComicRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{
			Message: "Body Invalido",
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

// GetComics godoc
// @Summary Mostra todos os quadrinhos do usuario
// @Description Mostra todos os quadrinhos do usuario
// @Tags Comic
// @Produce json
// @Security TokenAuth
// @Success 200 {object} models.Comics
// @Failure 500 {object} responses.CommonResponse
// @Failure 401 {object} responses.CommonResponse
// @Router /comic [get]
func (h *ComicHandler) GetComics(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	comics, err := h.svc.GetComics(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.CommonResponse{
			Message: "Erro ao buscar comics",
			Success: false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(comics)
}
