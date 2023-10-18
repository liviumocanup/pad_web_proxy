package services

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
	"track_service/config"
	"track_service/models"
	"track_service/repositories"
)

type TrackService interface {
	Upload(ctx context.Context, metadata models.TrackMetadata) (*models.TrackResponse, error)
	GetInfoById(ctx context.Context, request models.TrackIdRequest) (*models.TrackInfoResponse, error)
	EditInfo(ctx context.Context, request models.EditTrackRequest) error
	DeleteById(ctx context.Context, request models.TrackIdRequest) error
	FindAll(ctx context.Context) ([]*models.TrackInfoResponse, error)
	Status(ctx context.Context) (string, error)
}

type TrackServiceServer struct {
	repository repositories.TrackRepository
}

func NewTrackService(repository repositories.TrackRepository, cfg *config.Config) TrackService {
	log.Println("Creating track service")

	return &trackService{
		repository: repository,
		cfg:        cfg,
	}
}

type trackService struct {
	repository      repositories.TrackRepository
	cfg             *config.Config
	requestsCounter int64
	lastRequestTime time.Time
}

func (s *trackService) Upload(ctx context.Context, metadata models.TrackMetadata) (*models.TrackResponse, error) {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	////TEST FOR CONCURRENT TASK LIMIT
	//select {
	//case <-time.After(2 * time.Second):
	//	fmt.Println("Sleep Over.....")
	//case <-ctx.Done():
	//	return nil, ctx.Err()
	//}

	existingTrack, err := s.repository.GetByTitleAndUserId(metadata.Title, metadata.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.Internal, "internal error")
	}
	if existingTrack != nil {
		return nil, status.Error(codes.AlreadyExists, "track already exists")
	}

	track := &models.Track{
		Title:  metadata.Title,
		Artist: metadata.Artist,
		Album:  metadata.Album,
		Genre:  metadata.Genre,
		UserID: metadata.UserID,
		URL:    metadata.URL,
	}

	if err := s.repository.Create(track); err != nil {
		return nil, status.Error(codes.Internal, "failed to create track")
	}

	return &models.TrackResponse{
		TrackId: strconv.Itoa(int(track.ID)),
		URL:     track.URL,
		UserID:  track.UserID,
	}, nil
}

func (s *trackService) GetInfoById(ctx context.Context, request models.TrackIdRequest) (*models.TrackInfoResponse, error) {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	track, err := s.repository.GetById(request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "track not found")
	}

	return &models.TrackInfoResponse{
		TrackId: strconv.Itoa(int(track.ID)),
		Title:   track.Title,
		Artist:  track.Artist,
		Album:   track.Album,
		Genre:   track.Genre,
		URL:     track.URL,
		UserID:  track.UserID,
	}, nil
}

func (s *trackService) EditInfo(ctx context.Context, request models.EditTrackRequest) error {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	track, err := s.repository.GetById(request.TrackId)
	if err != nil {
		return status.Error(codes.NotFound, "track not found")
	}

	track.Title = request.Metadata.Title
	track.Artist = request.Metadata.Artist
	track.Album = request.Metadata.Album
	track.Genre = request.Metadata.Genre
	track.UserID = request.Metadata.UserID
	track.URL = request.Metadata.URL

	if err := s.repository.Update(track); err != nil {
		return status.Error(codes.Internal, "failed to update track info")
	}

	return nil
}

func (s *trackService) DeleteById(ctx context.Context, request models.TrackIdRequest) error {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	if err := s.repository.Delete(request.Id); err != nil {
		return status.Error(codes.Internal, "failed to delete track")
	}
	return nil
}

func (s *trackService) FindAll(ctx context.Context) ([]*models.TrackInfoResponse, error) {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	tracks, err := s.repository.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	var trackResponses []*models.TrackInfoResponse
	for _, track := range tracks {
		trackResponses = append(trackResponses, &models.TrackInfoResponse{
			TrackId: strconv.Itoa(int(track.ID)),
			Title:   track.Title,
			Artist:  track.Artist,
			Album:   track.Album,
			Genre:   track.Genre,
			URL:     track.URL,
			UserID:  track.UserID,
		})
	}

	return trackResponses, nil
}

func (s *trackService) Status(ctx context.Context) (string, error) {
	s.MonitorRequests()
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	return "ok", nil
}

func (s *trackService) MonitorRequests() {
	s.requestsCounter++
	currentTime := time.Now()
	if currentTime.Sub(s.lastRequestTime) >= time.Second {
		if s.requestsCounter >= int64(s.cfg.CriticalLoad) {
			log.Printf("ALERT: Critical load reached with %d requests in the last second!", s.requestsCounter)
		}
		s.requestsCounter = 0
		s.lastRequestTime = currentTime
	}
}
