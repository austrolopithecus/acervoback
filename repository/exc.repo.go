package repository

import (
	"acervoback/models"
	"gorm.io/gorm"
)

type ExchangeRepo interface {
	Create(exchange *models.Exchange) error
	FindByID(id string) (models.Exchange, error)
	FindByUser(userID string) ([]models.Exchange, error)
	Update(exchange *models.Exchange) error
}

type ExchangeRepoImpl struct {
	db *gorm.DB
}

func NewExchangeRepo(db *gorm.DB) ExchangeRepo {
	return &ExchangeRepoImpl{db: db}
}

func (e *ExchangeRepoImpl) Create(exchange *models.Exchange) error {
	return e.db.Create(exchange).Error
}

func (e *ExchangeRepoImpl) FindByID(id string) (models.Exchange, error) {
	var exchange models.Exchange
	err := e.db.First(&exchange, id).Error
	return exchange, err
}

func (e *ExchangeRepoImpl) FindByUser(userID string) ([]models.Exchange, error) {
	var exchanges []models.Exchange
	err := e.db.Where("requester_id = ? OR owner_id = ?", userID, userID).Find(&exchanges).Error
	return exchanges, err
}

func (e *ExchangeRepoImpl) Update(exchange *models.Exchange) error {
	return e.db.Save(exchange).Error
}
