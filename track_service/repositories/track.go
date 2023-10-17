package repositories

import (
	"gorm.io/gorm"
	"log"
	"track_service/models"
)

type TrackRepository interface {
	Create(track *models.Track) error
	GetById(id string) (*models.Track, error)
	Update(track *models.Track) error
	Delete(id string) error
	FindAll() ([]*models.Track, error)
	GetByUserId(userId string) ([]*models.Track, error)
	GetByTitleAndUserId(title string, userId string) (*models.Track, error)
}

func NewTrackRepository(db *gorm.DB) TrackRepository {
	log.Println("Creating new track repository")

	return &trackRepository{
		db: db,
	}
}

type trackRepository struct {
	db *gorm.DB
}

func (repo *trackRepository) Create(track *models.Track) error {
	return repo.db.Create(track).Error
}

func (repo *trackRepository) GetById(id string) (*models.Track, error) {
	var track models.Track
	err := repo.db.Where("id = ?", id).First(&track).Error
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (repo *trackRepository) Update(track *models.Track) error {
	return repo.db.Save(track).Error
}

func (repo *trackRepository) Delete(id string) error {
	return repo.db.Delete(&models.Track{}, "id = ?", id).Error
}

func (repo *trackRepository) FindAll() ([]*models.Track, error) {
	var users []*models.Track
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *trackRepository) GetByUserId(userId string) ([]*models.Track, error) {
	var tracks []*models.Track
	err := repo.db.Where("user_id = ?", userId).Find(&tracks).Error
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (repo *trackRepository) GetByTitleAndUserId(title string, userId string) (*models.Track, error) {
	var track models.Track
	err := repo.db.Where("title = ? AND user_id = ?", title, userId).First(&track).Error
	if err != nil {
		return nil, err
	}
	return &track, nil
}
