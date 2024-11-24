package repository

import (
	"acervoback/models"
	"github.com/rs/zerolog/log"
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
	log.Debug().Str("id", id).Msg("Buscando usuario por ID")
	var user models.User
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		log.Err(err).Msg("Usuario nao encontrado ou erro no banco")
		return models.User{}, err
	}
	log.Debug().Str("id", user.ID).Msg("Usuario encontrado")
	return user, nil
}

func (u *UserRepoImpl) Create(user models.User) (models.User, error) {
	log.Debug().Str("id", user.ID).Str("email", user.Email).Msg("Salvando usuario no banco")
	err := u.db.Create(&user).Error
	if err != nil {
		log.Err(err).Msg("Erro ao salvar usuario no banco")
		return models.User{}, err
	}
	log.Debug().Str("id", user.ID).Msg("Usuario salvo com sucesso")
	return user, nil
}

func (u *UserRepoImpl) Update(user models.User) (models.User, error) {
	return user, u.db.Save(user).Error
}

func (u *UserRepoImpl) Delete(id string) error {
	return u.db.Delete(&models.User{}, id).Error
}
