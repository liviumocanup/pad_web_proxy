package repositories

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"user_service/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	Delete(id string) error
	FindAll() ([]*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	log.Info().Msg("Creating new user repository")

	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Delete(id string) error {
	return repo.db.Delete(&models.User{}, "id = ?", id).Error
}
func (repo *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
