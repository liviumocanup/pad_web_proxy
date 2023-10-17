package services

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"playback_service/config"
	"playback_service/models"
	"playback_service/repositories"
	"strconv"
)

type PlaybackService interface {
	CreatePlaylist(ctx context.Context, request models.CreatePlaylistRequest) (*models.PlaylistResponse, error)
	RemovePlaylist(ctx context.Context, request models.RemovePlaylistRequest) error
	AddTracksToPlaylist(ctx context.Context, request models.AddTracksToPlaylistRequest) error
	RemoveTracksFromPlaylist(ctx context.Context, request models.RemoveTracksFromPlaylistRequest) error
	GetPlaylistById(ctx context.Context, request models.PlaylistIdRequest) (*models.PlaylistResponse, error)
	PlayPlaylist(ctx context.Context, request models.PlaylistIdRequest) (*models.PlayPlaylistResponse, error)
	Status(ctx context.Context) (string, error)
}

type PlaybackServiceServer struct {
	repository         repositories.PlaybackRepository
	trackServiceClient TrackServiceClient
}

func NewPlaybackService(repository repositories.PlaybackRepository, trackServiceClient TrackServiceClient, cfg *config.Config) PlaybackService {
	log.Println("Creating playback service")

	return &playbackService{
		repository:         repository,
		trackServiceClient: trackServiceClient,
		cfg:                cfg,
	}
}

type playbackService struct {
	repository         repositories.PlaybackRepository
	trackServiceClient TrackServiceClient
	cfg                *config.Config
}

func (s *playbackService) CreatePlaylist(ctx context.Context, request models.CreatePlaylistRequest) (*models.PlaylistResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	existingPlaylist, err := s.repository.GetByNameAndUserId(request.Name, request.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Error(codes.Internal, "internal error")
	}
	if existingPlaylist != nil {
		return nil, status.Error(codes.AlreadyExists, "playlist already exists")
	}

	playlist := &models.Playlist{
		Name:   request.Name,
		UserID: request.UserID,
	}

	if err := s.repository.Create(playlist); err != nil {
		return nil, status.Error(codes.Internal, "failed to create playlist")
	}

	return &models.PlaylistResponse{
		PlaylistId: strconv.Itoa(int(playlist.ID)),
		Name:       playlist.Name,
		Tracks:     []models.TrackMetadata{},
	}, nil
}

func (s *playbackService) RemovePlaylist(ctx context.Context, request models.RemovePlaylistRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	playlistIdUint, err := strconv.Atoi(request.PlaylistId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid playlist id")
	}

	playlist, err := s.repository.GetById(uint(playlistIdUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, "playlist not found")
		}
		return status.Error(codes.Internal, "failed to fetch playlist")
	}

	if playlist.UserID != request.UserId {
		return status.Error(codes.PermissionDenied, "you do not have the permission to remove this playlist")
	}

	if err := s.repository.RemoveById(uint(playlistIdUint)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, "playlist not found")
		}
		return status.Error(codes.Internal, "failed to remove playlist")
	}

	return nil
}

func (s *playbackService) AddTracksToPlaylist(ctx context.Context, request models.AddTracksToPlaylistRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	playlistIdUint, err := strconv.Atoi(request.PlaylistId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid playlist id")
	}

	playlist, err := s.repository.GetById(uint(playlistIdUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, "playlist not found")
		}
		return status.Error(codes.Internal, "failed to fetch playlist")
	}

	if playlist.UserID != request.UserId {
		return status.Error(codes.PermissionDenied, "you do not have the permission to add tracks to this playlist")
	}

	var playlistTracks []models.PlaylistTrack
	for _, trackIdStr := range request.TrackIds {
		_, err := strconv.Atoi(trackIdStr)
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid track id")
		}

		if isTrackInPlaylist(trackIdStr, playlist.Tracks) {
			continue
		}

		trackInfo, err := s.trackServiceClient.GetInfoById(ctx, trackIdStr)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		playlistTracks = append(playlistTracks, models.PlaylistTrack{
			TrackID: trackIdStr,
			TrackMetadata: models.TrackMetadata{
				Title:  trackInfo.Title,
				Artist: trackInfo.Artist,
				Album:  trackInfo.Album,
				Genre:  trackInfo.Genre,
				UserID: trackInfo.UserId,
				URL:    trackInfo.Url,
			},
		})
	}

	playlist.Tracks = append(playlist.Tracks, playlistTracks...)

	if err := s.repository.Update(playlist); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (s *playbackService) RemoveTracksFromPlaylist(ctx context.Context, request models.RemoveTracksFromPlaylistRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	playlistIdUint, err := strconv.Atoi(request.PlaylistId)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid playlist id")
	}

	playlist, err := s.repository.GetById(uint(playlistIdUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, "playlist not found")
		}
		return status.Error(codes.Internal, "failed to fetch playlist")
	}

	if playlist.UserID != request.UserId {
		return status.Error(codes.PermissionDenied, "you do not have the permission to add tracks to this playlist")
	}

	var trackIDsToRemove []string
	for _, trackIdStr := range request.TrackIds {
		_, err := strconv.Atoi(trackIdStr)
		if err != nil {
			return status.Error(codes.InvalidArgument, "invalid track id")
		}
		trackIDsToRemove = append(trackIDsToRemove, trackIdStr)
	}

	if err := s.repository.DeleteTracks(uint(playlistIdUint), trackIDsToRemove); err != nil {
		return status.Error(codes.Internal, "failed to remove tracks from playlist")
	}

	return nil
}

func (s *playbackService) GetPlaylistById(ctx context.Context, request models.PlaylistIdRequest) (*models.PlaylistResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	// TEST FOR CONCURRENT TASK LIMIT
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("Sleep Over.....")
	//case <-ctx.Done():
	//	return nil, ctx.Err()
	//}

	playlistIdUint, err := strconv.Atoi(request.PlaylistId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid playlist id")
	}

	playlist, err := s.repository.GetById(uint(playlistIdUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "playlist not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	var trackMetadatas []models.TrackMetadata
	for _, track := range playlist.Tracks {
		trackMetadatas = append(trackMetadatas, track.TrackMetadata)
	}

	return &models.PlaylistResponse{
		PlaylistId: strconv.Itoa(int(playlist.ID)),
		Name:       playlist.Name,
		Tracks:     trackMetadatas,
	}, nil
}

func (s *playbackService) PlayPlaylist(ctx context.Context, request models.PlaylistIdRequest) (*models.PlayPlaylistResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	playlistIdUint, err := strconv.Atoi(request.PlaylistId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid playlist id")
	}

	playlist, err := s.repository.GetById(uint(playlistIdUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "playlist not found")
		}
		return nil, status.Error(codes.Internal, "failed to fetch playlist")
	}

	var trackPlayMetadatas []models.TrackPlayMetadata
	for _, track := range playlist.Tracks {
		trackPlayMetadatas = append(trackPlayMetadatas, models.TrackPlayMetadata{
			Id:    track.TrackID,
			Title: track.Title,
			URL:   track.URL,
		})
	}

	return &models.PlayPlaylistResponse{
		PlaylistName: playlist.Name,
		Tracks:       trackPlayMetadatas,
	}, nil
}

func isTrackInPlaylist(trackID string, tracks []models.PlaylistTrack) bool {
	for _, track := range tracks {
		if track.TrackID == trackID {
			return true
		}
	}
	return false
}

func (s *playbackService) Status(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	return "ok", nil
}
