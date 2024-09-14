package repository

import (
    "acervoback/models"
    "gorm.io/gorm"
)

// Defina a interface para Reviews
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

func (c *ComicRepoImpl) FindByID(id string) (models.Comic, error) {
    var comic models.Comic
    err := c.db.First(&comic, "id = ?", id).Error
    return comic, err
}


func (r *ReviewRepoImpl) Update(review *models.Review) error {
    return r.db.Save(review).Error
}

