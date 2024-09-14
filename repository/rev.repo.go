package repository

import (
	"acervoback/models"
	"gorm.io/gorm"
)

type ReviewRepo interface {
	Create(review *models.Review) error
	FindByComicID(comicID string) ([]models.Review, error)
}

type ReviewRepoImpl struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) ReviewRepo {
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

