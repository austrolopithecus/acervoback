package repository

import (
	"acervoback/models"
	"gorm.io/gorm"
)

type ComicRepo interface {
	Create(comic *models.Comic) error
	FindAll() ([]models.Comic, error)
	FindByID(id int) (models.Comic, error)
	Update(comic *models.Comic) error
	Delete(id string) error
	FindByOwner(ownerID string) ([]models.Comic, error)
}

type ComicRepoImpl struct {
	db *gorm.DB
}

func (c *ComicRepoImpl) Create(comic *models.Comic) error {
	return c.db.Create(comic).Error
}

func (c *ComicRepoImpl) FindAll() ([]models.Comic, error) {
	var comics []models.Comic
	err := c.db.Find(&comics).Error
	return comics, err
}

func (c *ComicRepoImpl) FindByID(id int) (models.Comic, error) {
	var comic models.Comic
	err := c.db.First(&comic, id).Error
	return comic, err
}

func (c *ComicRepoImpl) Update(comic *models.Comic) error {
	return c.db.Save(comic).Error
}

func (c *ComicRepoImpl) Delete(id string) error {
	return c.db.Delete(&models.Comic{}, id).Error
}

func (c *ComicRepoImpl) FindByOwner(ownerID string) ([]models.Comic, error) {
	var comics []models.Comic
	err := c.db.Where("user_id = ?", ownerID).Find(&comics).Error
	return comics, err
}

func NewComicRepoImpl(db *gorm.DB) ComicRepo {
	return &ComicRepoImpl{db: db}
}
