package services

import (
	"acervoback/models"
	"acervoback/repository"
	"errors"
	"github.com/google/uuid"
)

type ExchangeService struct {
	exchangeRepo repository.ExchangeRepo
	comicRepo    repository.ComicRepo
	userRepo     repository.UserRepo
}

func NewExchangeService(exchangeRepo repository.ExchangeRepo, comicRepo repository.ComicRepo, userRepo repository.UserRepo) *ExchangeService {
	return &ExchangeService{exchangeRepo: exchangeRepo, comicRepo: comicRepo, userRepo: userRepo}
}

func (s *ExchangeService) RequestExchange(user1ID, user2ID, comic1ID, comic2ID string) (*models.Exchange, error) {
	// Verificar se os quadrinhos existem e pertencem aos usuários corretos
	comic1, err := s.comicRepo.FindByID(comic1ID)
	if err != nil || comic1.UserID != user1ID {
		return nil, errors.New("invalid comic for user1")
	}

	comic2, err := s.comicRepo.FindByID(comic2ID)
	if err != nil || comic2.UserID != user2ID {
		return nil, errors.New("invalid comic for user2")
	}

	// Criar uma nova solicitação de troca
	exchange := &models.Exchange{
		ID:       uuid.New().String(),
		User1ID:  user1ID,
		User2ID:  user2ID,
		Comic1ID: comic1ID,
		Comic2ID: comic2ID,
		Status:   "pending",
	}

	err = s.exchangeRepo.Create(exchange)
	if err != nil {
		return nil, err
	}

	return exchange, nil
}

func (s *ExchangeService) CompleteExchange(exchangeID string) error {
	exchange, err := s.exchangeRepo.FindByID(exchangeID)
	if err != nil {
		return err
	}

	if exchange.Status != "pending" {
		return errors.New("exchange is not in pending state")
	}

	// Trocar a propriedade dos quadrinhos
	comic1, err := s.comicRepo.FindByID(exchange.Comic1ID)
	if err != nil {
		return err
	}
	comic1.UserID = exchange.User2ID
	err = s.comicRepo.Update(&comic1)
	if err != nil {
		return err
	}

	comic2, err := s.comicRepo.FindByID(exchange.Comic2ID)
	if err != nil {
		return err
	}
	comic2.UserID = exchange.User1ID
	err = s.comicRepo.Update(&comic2)
	if err != nil {
		return err
	}

	// Atualizar o status da troca
	exchange.Status = "completed"
	return s.exchangeRepo.Update(&exchange)
}

