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
	"track_service/config"
	"track_service/models"
	"track_service/proto"
	commonpb "track_service/proto/common"
	"track_service/services"
)

func NewGrpcServer(cfg *config.Config, trackService services.TrackService) (*grpc.Server, net.Listener, error) {
	log.Println("Creating new gRPC server for track service")

	server := &grpcServer{
		trackService: trackService,
		semaphore:    make(chan struct{}, cfg.ConcurrentLimit),
	}

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, nil, err
	}

	srv := grpc.NewServer()
	proto.RegisterTrackServiceServer(srv, server)

	return srv, listener, nil
}

type grpcServer struct {
	proto.UnsafeTrackServiceServer
	trackService services.TrackService
	semaphore    chan struct{}
}

func (s *grpcServer) Upload(ctx context.Context, req *commonpb.TrackMetadata) (*proto.TrackResponse, error) {
	s.acquire()
	defer s.release()

	response, err := s.trackService.Upload(ctx, models.TrackMetadata{
		Title:  req.Title,
		Artist: req.Artist,
		Album:  req.Album,
		Genre:  req.Genre,
		UserID: req.UserId,
		URL:    req.Url,
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.TrackResponse{
		TrackId: response.TrackId,
		Url:     response.URL,
		UserId:  response.UserID,
	}, nil
}

func (s *grpcServer) GetInfoById(ctx context.Context, req *proto.TrackIdRequest) (*proto.TrackInfoResponse, error) {
	s.acquire()
	defer s.release()

	response, err := s.trackService.GetInfoById(ctx, models.TrackIdRequest{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.TrackInfoResponse{
		TrackId: response.TrackId,
		Title:   response.Title,
		Artist:  response.Artist,
		Album:   response.Album,
		Genre:   response.Genre,
		Url:     response.URL,
		UserId:  response.UserID,
	}, nil
}

func (s *grpcServer) EditInfo(ctx context.Context, req *proto.EditTrackRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	metadata := models.TrackMetadata{
		Title:  req.Metadata.Title,
		Artist: req.Metadata.Artist,
		Album:  req.Metadata.Album,
		Genre:  req.Metadata.Genre,
		UserID: req.Metadata.UserId,
		URL:    req.Metadata.Url,
	}
	err := s.trackService.EditInfo(ctx, models.EditTrackRequest{
		TrackId:  req.TrackId,
		Metadata: metadata,
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) DeleteById(ctx context.Context, req *proto.TrackIdRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	err := s.trackService.DeleteById(ctx, models.TrackIdRequest{
		Id: req.Id,
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) FindAll(ctx context.Context, _ *empty.Empty) (*proto.TrackListResponse, error) {
	s.acquire()
	defer s.release()

	tracks, err := s.trackService.FindAll(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	var protoTracks []*proto.TrackInfoResponse
	for _, track := range tracks {
		protoTrack := &proto.TrackInfoResponse{
			TrackId: track.TrackId,
			Title:   track.Title,
			Artist:  track.Artist,
			Album:   track.Album,
			Genre:   track.Genre,
			Url:     track.URL,
			UserId:  track.UserID,
		}
		protoTracks = append(protoTracks, protoTrack)
	}

	return &proto.TrackListResponse{
		Tracks: protoTracks,
	}, nil
}

func (s *grpcServer) Status(ctx context.Context, _ *empty.Empty) (*commonpb.StatusResponse, error) {
	s.acquire()
	defer s.release()

	serviceStatus, err := s.trackService.Status(ctx)
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
