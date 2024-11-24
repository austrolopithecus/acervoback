package services

import (
	"acervoback/models"
	"acervoback/models/requests"
	"acervoback/repository"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type UserService struct {
	userRepo repository.UserRepo
	jwtRepo  repository.JWTRepo
}

func NewUserService(userRepo repository.UserRepo, jwtRepo repository.JWTRepo) *UserService {
	return &UserService{userRepo: userRepo, jwtRepo: jwtRepo}
}

func (u *UserService) Register(body requests.UserRegisterRequest) (string, error) {
	user := models.User{
		ID:       uuid.New().String(),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}
	log.Debug().Str("id", user.ID).Str("name", user.Name).Str("email", user.Email).Msg("Criando usuario")

	_, err := u.userRepo.Create(user)
	if err != nil {
		log.Err(err).Msg("Erro ao criar usuario")
		return "", err
	}
	return u.jwtRepo.GenerateToken(user.ID)
}

func (u *UserService) Login(request requests.UserLoginRequest) (string, error) {
	log.Trace().Str("email", request.Email).Msg("Buscando usuario")
	user, err := u.userRepo.FindByEmail(request.Email)
	if err != nil {
		log.Error().Err(err).Msg("Erro ao buscar usuario")
		return "", err
	}
	if user.Password != request.Password {
		log.Debug().Msg("Senha incorreta")
		return "", nil
	}
	return u.jwtRepo.GenerateToken(user.ID)
}

func (u *UserService) Me(userID string) (models.User, error) {
	log.Trace().Str("id", userID).Msg("Buscando usuario")
	return u.userRepo.FindByID(userID)
}

func (u *UserService) Jwt(token string) (string, error) {
	if u.jwtRepo == nil {
		log.Error().Msg("jwtRepo not initialized")
		return "", fmt.Errorf("jwtRepo not initialized")
	}
	return u.jwtRepo.VerifyToken(token)
}
