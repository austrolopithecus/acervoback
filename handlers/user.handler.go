package handlers

import (
	"acervoback/models/requests"
	"acervoback/models/responses"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// UserRegister godoc
// @Summary Registra um novo usuário
// @Description Registra um novo usuário
// @Tags User
// @Accept json
// @Produce json
// @Param body body requests.UserRegisterRequest true "Body"
// @Success 200 {object} responses.UserLoginResponse
// @Failure 400 {object} responses.CommonResponse
// @Failure 409 {object} responses.CommonResponse
// @Failure 500 {object} responses.CommonResponse
// @Router /user/register [post]
func (u *UserHandler) Register(c *fiber.Ctx) error {
	var body requests.UserRegisterRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{
			Message: "Invalid request",
			Success: false,
		})
	}

	token, err := u.svc.Register(body)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") || strings.Contains(err.Error(), "duplicar valor da chave viola a restrição de unicidade") {
			return c.Status(fiber.StatusConflict).JSON(responses.CommonResponse{
				Message: "Email already exists",
				Success: false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(responses.CommonResponse{
			Message: "Internal server error",
			Success: false,
		})
	}
	if token == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.CommonResponse{
			Message: "Internal server error",
			Success: false})
	}
	return c.JSON(responses.UserLoginResponse{
		CommonResponse: responses.CommonResponse{
			Message: "User registered successfully",
			Success: true,
		},
		Token: token,
	})
}

// UserLogin godoc
// @Summary Login de  um  usuário
// @Description Login de  um  usuário
// @Tags User
// @Accept json
// @Produce json
// @Param body body requests.UserLoginRequest true "Body"
// @Success 200 {object} responses.UserLoginResponse
// @Failure 400 {object} responses.CommonResponse
// @Failure 401 {object} responses.CommonResponse
// @Failure 500 {object} responses.CommonResponse
// @Router /user/login [post]
func (u *UserHandler) Login(c *fiber.Ctx) error {
	var body requests.UserLoginRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{
			Message: "Invalid request",
			Success: false,
		})
	}
	user, err := u.svc.Login(body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.CommonResponse{
			Message: "Invalid credentials",
			Success: false,
		})
	}
	if user == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.CommonResponse{
			Message: "Invalid credentials",
			Success: false,
		})
	}

	return c.JSON(responses.UserLoginResponse{
		CommonResponse: responses.CommonResponse{
			Message: "User logged in successfully",
			Success: true,
		},
		Token: user,
	})
}

// UserMe godoc
// @Summary Retorna os dados do usuário logado
// @Description Retorna os dados do usuário logado
// @Tags User
// @Produce json
// @Security TokenAuth
// @Success 200 {object} responses.UserMeResponse
// @Failure 500 {object} responses.CommonResponse
// @Failure 401 {object} responses.CommonResponse
// @Router /user/me [get]
func (u *UserHandler) Me(c *fiber.Ctx) error {
	id := c.Locals("userID").(string)
	user, err := u.svc.Me(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.CommonResponse{
			Message: "Internal server error",
			Success: false,
		})
	}
	return c.JSON(responses.UserMeResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
