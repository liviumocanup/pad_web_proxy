package adapter

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"playback_service/config"
	"playback_service/models"
	"playback_service/proto"
	commonpb "playback_service/proto/common"
	"playback_service/services"
)

func NewGrpcServer(cfg *config.Config, playbackService services.PlaybackService) (*grpc.Server, net.Listener, error) {
	log.Println("Creating new gRPC server for playback service")

	server := &grpcServer{
		playbackService: playbackService,
		semaphore:       make(chan struct{}, cfg.ConcurrentLimit),
	}

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, nil, err
	}

	srv := grpc.NewServer()
	proto.RegisterPlaybackServiceServer(srv, server)

	return srv, listener, nil
}

type grpcServer struct {
	proto.UnsafePlaybackServiceServer
	playbackService services.PlaybackService
	semaphore       chan struct{}
}

func (s *grpcServer) CreatePlaylist(ctx context.Context, req *proto.CreatePlaylistRequest) (*proto.PlaylistResponse, error) {
	s.acquire()
	defer s.release()

	response, err := s.playbackService.CreatePlaylist(ctx, models.CreatePlaylistRequest{
		Name:   req.Name,
		UserID: req.UserId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.PlaylistResponse{
		PlaylistId: response.PlaylistId,
		Name:       response.Name,
		Tracks:     convertToProtoTrackMetadata(response.Tracks),
	}, nil
}

func (s *grpcServer) RemovePlaylist(ctx context.Context, req *proto.RemovePlaylistRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	err := s.playbackService.RemovePlaylist(ctx, models.RemovePlaylistRequest{
		PlaylistId: req.PlaylistId,
		UserId:     req.UserId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) AddTracksToPlaylist(ctx context.Context, req *proto.AddTracksToPlaylistRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	err := s.playbackService.AddTracksToPlaylist(ctx, models.AddTracksToPlaylistRequest{
		PlaylistId: req.PlaylistId,
		TrackIds:   req.TrackIds,
		UserId:     req.UserId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) RemoveTracksFromPlaylist(ctx context.Context, req *proto.RemoveTracksFromPlaylistRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	err := s.playbackService.RemoveTracksFromPlaylist(ctx, models.RemoveTracksFromPlaylistRequest{
		PlaylistId: req.PlaylistId,
		TrackIds:   req.TrackIds,
		UserId:     req.UserId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) GetPlaylistById(ctx context.Context, req *proto.PlaylistIdRequest) (*proto.PlaylistResponse, error) {
	s.acquire()
	defer s.release()

	response, err := s.playbackService.GetPlaylistById(ctx, models.PlaylistIdRequest{
		PlaylistId: req.PlaylistId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.PlaylistResponse{
		PlaylistId: response.PlaylistId,
		Name:       response.Name,
		Tracks:     convertToProtoTrackMetadata(response.Tracks),
	}, nil
}

func (s *grpcServer) PlayPlaylist(ctx context.Context, req *proto.PlaylistIdRequest) (*proto.PlayPlaylistResponse, error) {
	s.acquire()
	defer s.release()

	response, err := s.playbackService.PlayPlaylist(ctx, models.PlaylistIdRequest{
		PlaylistId: req.PlaylistId,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.PlayPlaylistResponse{
		PlaylistName: response.PlaylistName,
		Tracks:       convertToProtoTrackPlayMetadata(response.Tracks),
	}, nil
}

func (s *grpcServer) Status(ctx context.Context, _ *empty.Empty) (*commonpb.StatusResponse, error) {
	s.acquire()
	defer s.release()

	serviceStatus, err := s.playbackService.Status(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &commonpb.StatusResponse{Status: serviceStatus}, nil
}

func (s *grpcServer) acquire() {
	s.semaphore <- struct{}{}
}

func (s *grpcServer) release() {
	<-s.semaphore
}

func convertToProtoTrackMetadata(tracks []models.TrackMetadata) []*commonpb.TrackMetadata {
	var protoTracks []*commonpb.TrackMetadata
	for _, t := range tracks {
		protoTracks = append(protoTracks, &commonpb.TrackMetadata{
			Title:  t.Title,
			Artist: t.Artist,
			Album:  t.Album,
			Genre:  t.Genre,
			UserId: t.UserID,
			Url:    t.URL,
		})
	}
	return protoTracks
}

func convertToProtoTrackPlayMetadata(tracks []models.TrackPlayMetadata) []*proto.TrackPlayMetadata {
	var protoTracks []*proto.TrackPlayMetadata
	for _, t := range tracks {
		protoTracks = append(protoTracks, &proto.TrackPlayMetadata{
			TrackId: t.Id,
			Title:   t.Title,
			Url:     t.URL,
		})
	}
	return protoTracks
}
