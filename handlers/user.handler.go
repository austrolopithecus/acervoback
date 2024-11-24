package handlers

import (
	"acervoback/models/requests"
	"acervoback/models/responses"
	"acervoback/services"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strings"
)

type UserHandler struct {
	svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Middleware JWT para autenticação
func (u *UserHandler) JwtMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	userID, err := u.svc.Jwt(auth)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	c.Locals("userID", userID)
	return c.Next()
}

// Funções relacionadas ao usuário (registro, login, perfil)
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
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
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

	return c.JSON(responses.UserLoginResponse{
		CommonResponse: responses.CommonResponse{
			Message: "User registered successfully",
			Success: true,
		},
		Token: token,
	})
}

// Função de login
func (u *UserHandler) Login(c *fiber.Ctx) error {
	var body requests.UserLoginRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.CommonResponse{
			Message: "Invalid request",
			Success: false,
		})
	}

	token, err := u.svc.Login(body)
	if err != nil {
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
		Token: token,
	})
}

// Função para obter os dados do usuário
func (u *UserHandler) Me(c *fiber.Ctx) error {
	id := c.Locals("userID").(string)
	user, err := u.svc.Me(id)
	if err != nil {
		log.Err(err).Msg("Erro ao buscar perfil do usuario")
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
