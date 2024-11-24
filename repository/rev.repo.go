package repository

import (
	"acervoback/models"
	"gorm.io/gorm"
)

// Interface para o repositório de reviews
type ReviewRepo interface {
	Create(review *models.Review) error
	FindByComicID(comicID string) ([]models.Review, error)
	FindByID(id string) (models.Review, error)
	Update(review *models.Review) error
}

// Implementação concreta do ReviewRepo
type ReviewRepoImpl struct {
	db *gorm.DB
}

// Função para criar uma nova instância do repositório
func NewReviewRepoImpl(db *gorm.DB) ReviewRepo {
	return &ReviewRepoImpl{db: db}
}

func (r *ReviewRepoImpl) Create(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepoImpl) FindByComicID(comicID string) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("comic_id = ?", comicID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepoImpl) FindByID(id string) (models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	return review, err
}

func (r *ReviewRepoImpl) Update(review *models.Review) error {
	return r.db.Save(review).Error
}

