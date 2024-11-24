package services

import (
	"acervoback/models"
	"acervoback/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ExchangeService struct {
	exchangeRepo repository.ExchangeRepo
	comicRepo    repository.ComicRepo
}

func NewExchangeService(exchangeRepo repository.ExchangeRepo, comicRepo repository.ComicRepo) *ExchangeService {
	return &ExchangeService{
		exchangeRepo: exchangeRepo,
		comicRepo:    comicRepo,
	}
}

func (s *ExchangeService) RequestExchange(comicID, requesterID, ownerID string) (*models.Exchange, error) {
	// Valida se o quadrinho pertence ao proprietário
	comic, err := s.comicRepo.FindByID(comicID)
	if err != nil || comic.UserID != ownerID {
		return nil, errors.New("invalid comic or ownership")
	}

	exchange := &models.Exchange{
		ID:          uuid.New().String(),
		ComicID:     comicID,
		RequesterID: requesterID,
		OwnerID:     ownerID,
		Status:      models.Pending,
	}

	err = s.exchangeRepo.Create(exchange)
	if err != nil {
		return nil, err
	}

	return exchange, nil
}

func (s *ExchangeService) AcceptExchange(exchangeID string, userID string) error {
	exchange, err := s.exchangeRepo.FindByID(exchangeID)
	if err != nil {
		return fmt.Errorf("Exchange não encontrado: %w", err)
	}

	if exchange.OwnerID != userID {
		return fmt.Errorf("Apenas o proprietário pode aceitar a troca")
	}

	exchange.Status = models.Accepted
	return s.exchangeRepo.Update(&exchange)
}

func (s *ExchangeService) DeclineExchange(exchangeID, userID string) error {
	exchange, err := s.exchangeRepo.FindByID(exchangeID)
	if err != nil {
		return err
	}

	if exchange.OwnerID != userID || exchange.Status != models.Pending {
		return errors.New("exchange not allowed")
	}

	exchange.Status = models.Declined
	return s.exchangeRepo.Update(&exchange)
}
