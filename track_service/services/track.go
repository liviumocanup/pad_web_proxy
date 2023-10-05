package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"track_service/models"
	"track_service/repositories"
)

type TrackService interface {
	Upload(ctx context.Context, metadata models.TrackMetadata) (*models.TrackResponse, error)
	GetInfoById(ctx context.Context, request models.IdRequest) (*models.TrackInfoResponse, error)
	EditInfo(ctx context.Context, request models.EditTrackRequest) error
	DeleteById(ctx context.Context, request models.IdRequest) error
}

type TrackServiceServer struct {
	repository repositories.TrackRepository
}

func NewTrackService(repository repositories.TrackRepository) TrackService {
	log.Println("Creating track service")

	return &trackService{
		repository: repository,
	}
}

type trackService struct {
	repository repositories.TrackRepository
}

func (s *trackService) Upload(ctx context.Context, metadata models.TrackMetadata) (*models.TrackResponse, error) {
	// Logic for S3 upload can be added here
	// For now, just adding to DB.

	track := &models.Track{
		Title:  metadata.Title,
		Artist: metadata.Artist,
		Album:  metadata.Album,
		Genre:  metadata.Genre,
		// URL can be updated with S3 URL once uploaded
	}

	if err := s.repository.Create(track); err != nil {
		return nil, status.Error(codes.Internal, "failed to create track")
	}

	return &models.TrackResponse{
		TrackId: strconv.Itoa(int(track.ID)),
		URL:     track.URL,
	}, nil
}

func (s *trackService) GetInfoById(ctx context.Context, request models.IdRequest) (*models.TrackInfoResponse, error) {
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
	}, nil
}

func (s *trackService) EditInfo(ctx context.Context, request models.EditTrackRequest) error {
	track, err := s.repository.GetById(request.TrackId)
	if err != nil {
		return status.Error(codes.NotFound, "track not found")
	}

	track.Title = request.Metadata.Title
	track.Artist = request.Metadata.Artist
	track.Album = request.Metadata.Album
	track.Genre = request.Metadata.Genre
	// Update URL if needed.

	if err := s.repository.Update(track); err != nil {
		return status.Error(codes.Internal, "failed to update track info")
	}

	return nil
}

func (s *trackService) DeleteById(ctx context.Context, request models.IdRequest) error {
	if err := s.repository.Delete(request.Id); err != nil {
		return status.Error(codes.Internal, "failed to delete track")
	}
	return nil
}
