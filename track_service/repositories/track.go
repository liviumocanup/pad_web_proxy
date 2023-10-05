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
