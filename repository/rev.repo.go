package repository

import (
    "acervoback/models"
    "gorm.io/gorm"
)

type ExchangeRepo interface {
    Create(exchange *models.Exchange) error
    FindByID(id string) (models.Exchange, error)
    Update(exchange *models.Exchange) error
}

type ExchangeRepoImpl struct {
    db *gorm.DB
}

func NewExchangeRepoImpl(db *gorm.DB) ExchangeRepo {
    return &ExchangeRepoImpl{db: db}
}

func (r *ExchangeRepoImpl) Create(exchange *models.Exchange) error {
    return r.db.Create(exchange).Error
}

func (r *ExchangeRepoImpl) FindByID(id string) (models.Exchange, error) {
    var exchange models.Exchange
    err := r.db.First(&exchange, "id = ?", id).Error
    return exchange, err
}

func (r *ExchangeRepoImpl) Update(exchange *models.Exchange) error {
    return r.db.Save(exchange).Error
}

