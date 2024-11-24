package repository

import (
	"acervoback/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type ComicRepo interface {
	Create(comic *models.Comic) error
	FindAll() ([]models.Comic, error)
	FindByID(id string) (models.Comic, error)
	Update(comic *models.Comic) error
	Delete(id string) error
	FindByOwner(ownerID string) ([]models.Comic, error)
	UpdateOwner(comicID, newOwnerID string) error
}

type ComicRepoImpl struct {
	db *gorm.DB
}

func (c *ComicRepoImpl) Create(comic *models.Comic) error {
	err := c.db.Create(comic).Error
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return fmt.Errorf("ISBN já existe")
	}
	return err
}

func (c *ComicRepoImpl) FindAll() ([]models.Comic, error) {
	var comics []models.Comic
	err := c.db.Find(&comics).Error
	return comics, err
}

func (c *ComicRepoImpl) FindByID(id string) (models.Comic, error) {
	var comic models.Comic
	err := c.db.First(&comic, "id = ?", id).Error
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

func (c *ComicRepoImpl) UpdateOwner(comicID, newOwnerID string) error {
	return c.db.Model(&models.Comic{}).Where("id = ?", comicID).Update("user_id", newOwnerID).Error
}

func NewComicRepoImpl(db *gorm.DB) ComicRepo {
	return &ComicRepoImpl{db: db}
}
