package repository

import (
	"acervoback/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindAll() ([]models.User, error)
	FindByID(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string) error
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (u *UserRepoImpl) FindAll() ([]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *UserRepoImpl) FindByID(id string) (models.User, error) {
	var user models.User
	result := u.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (u *UserRepoImpl) Create(user models.User) (models.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (u *UserRepoImpl) Update(user models.User) (models.User, error) {
	return user, u.db.Save(user).Error
}

func (u *UserRepoImpl) Delete(id string) error {
	return u.db.Delete(&models.User{}, id).Error
}
