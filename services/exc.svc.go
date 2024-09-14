// services/exchange.svc.go
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
}

func NewExchangeService(exchangeRepo repository.ExchangeRepo, comicRepo repository.ComicRepo) *ExchangeService {
    return &ExchangeService{
        exchangeRepo: exchangeRepo,
        comicRepo:    comicRepo,
    }
}

func (s *ExchangeService) RequestExchange(comicIDFrom, comicIDTo, userIDFrom, userIDTo string) (*models.Exchange, error) {
    // Verifica se os quadrinhos existem e pertencem aos usuários
    comicFrom, err := s.comicRepo.FindByID(comicIDFrom)
    if err != nil || comicFrom.UserID != userIDFrom {
        return nil, errors.New("invalid comic or ownership")
    }

    comicTo, err := s.comicRepo.FindByID(comicIDTo)
    if err != nil || comicTo.UserID != userIDTo {
        return nil, errors.New("invalid comic or ownership")
    }

    exchange := &models.Exchange{
        ID:          uuid.New().String(),
        ComicIDFrom: comicIDFrom,
        ComicIDTo:   comicIDTo,
        UserIDFrom:  userIDFrom,
        UserIDTo:    userIDTo,
        Status:      models.Pending,
    }

    err = s.exchangeRepo.Create(exchange)
    if err != nil {
        return nil, err
    }

    return exchange, nil
}

func (s *ExchangeService) AcceptExchange(exchangeID, userID string) error {
    exchange, err := s.exchangeRepo.FindByID(exchangeID)
    if err != nil {
        return err
    }

    if exchange.UserIDTo != userID || exchange.Status != models.Pending {
        return errors.New("exchange not allowed")
    }

    exchange.Status = models.Accepted
    return s.exchangeRepo.Update(&exchange)
}

func (s *ExchangeService) CompleteExchange(exchangeID, userID string) error {
    exchange, err := s.exchangeRepo.FindByID(exchangeID)
    if err != nil {
        return err
    }

    if exchange.UserIDFrom != userID || exchange.Status != models.Accepted {
        return errors.New("exchange not allowed")
    }

    // Trocar os quadrinhos
    err = s.comicRepo.UpdateOwner(exchange.ComicIDFrom, exchange.UserIDTo)
    if err != nil {
        return err
    }

    err = s.comicRepo.UpdateOwner(exchange.ComicIDTo, exchange.UserIDFrom)
    if err != nil {
        return err
    }

    exchange.Status = models.Completed
    return s.exchangeRepo.Update(&exchange)
}

