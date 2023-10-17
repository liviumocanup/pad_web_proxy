package repositories

import (
	"gorm.io/gorm"
	"log"
	"playback_service/models"
)

type PlaybackRepository interface {
	Create(playlist *models.Playlist) error
	RemoveById(id uint) error
	Update(playlist *models.Playlist) error
	GetById(id uint) (*models.Playlist, error)
	GetByNameAndUserId(name string, userId string) (*models.Playlist, error)
	DeleteTracks(playlistID uint, trackIDs []string) error
}

func NewPlaybackRepository(db *gorm.DB) PlaybackRepository {
	log.Println("Creating new playlist repository")

	return &playbackRepository{
		db: db,
	}
}

type playbackRepository struct {
	db *gorm.DB
}

func (repo *playbackRepository) Create(playlist *models.Playlist) error {
	return repo.db.Create(playlist).Error
}

func (repo *playbackRepository) GetByNameAndUserId(name string, userId string) (*models.Playlist, error) {
	var playlist models.Playlist
	err := repo.db.Preload("Tracks").Where("name = ? AND user_id = ?", name, userId).First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (repo *playbackRepository) RemoveById(id uint) error {
	// Begin a transaction
	tx := repo.db.Begin()

	// Delete tracks associated with the playlist from the playlist_tracks table
	if err := tx.Where("playlist_id = ?", id).Delete(&models.PlaylistTrack{}).Error; err != nil {
		tx.Rollback() // Rollback if there's an error
		return err
	}

	// Then delete the playlist
	if err := tx.Delete(&models.Playlist{}, "id = ?", id).Error; err != nil {
		tx.Rollback() // Rollback if there's an error
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (repo *playbackRepository) Update(playlist *models.Playlist) error {
	return repo.db.Save(playlist).Error
}

func (repo *playbackRepository) DeleteTracks(playlistID uint, trackIDs []string) error {
	return repo.db.Where("playlist_id = ? AND track_id IN ?", playlistID, trackIDs).Delete(&models.PlaylistTrack{}).Error
}

func (repo *playbackRepository) GetById(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	err := repo.db.Preload("Tracks").Where("id = ?", id).First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}
