package adapter

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	"track_service/config"
	"track_service/models"
	"track_service/proto"
	"track_service/services"
)

func NewGrpcServer(cfg *config.Config, trackService services.TrackService) (*grpc.Server, net.Listener, error) {
	log.Println("Creating new gRPC server for track service")

	server := &grpcServer{
		trackService: trackService,
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
}

func (s *grpcServer) Upload(ctx context.Context, req *proto.TrackMetadata) (*proto.TrackResponse, error) {
	response, err := s.trackService.Upload(ctx, models.TrackMetadata{
		Title:  req.Title,
		Artist: req.Artist,
		Album:  req.Album,
		Genre:  req.Genre,
	})

	if err != nil {
		return nil, err
	}

	return &proto.TrackResponse{
		TrackId: response.TrackId,
		Url:     response.URL,
	}, nil
}

func (s *grpcServer) GetInfoById(ctx context.Context, req *proto.IdRequest) (*proto.TrackInfoResponse, error) {
	response, err := s.trackService.GetInfoById(ctx, models.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &proto.TrackInfoResponse{
		TrackId: response.TrackId,
		Title:   response.Title,
		Artist:  response.Artist,
		Album:   response.Album,
		Genre:   response.Genre,
		Url:     response.URL,
	}, nil
}

func (s *grpcServer) EditInfo(ctx context.Context, req *proto.EditTrackRequest) (*empty.Empty, error) {
	metadata := models.TrackMetadata{
		Title:  req.Metadata.Title,
		Artist: req.Metadata.Artist,
		Album:  req.Metadata.Album,
		Genre:  req.Metadata.Genre,
	}
	err := s.trackService.EditInfo(ctx, models.EditTrackRequest{
		TrackId:  req.TrackId,
		Metadata: metadata,
	})
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) DeleteById(ctx context.Context, req *proto.IdRequest) (*empty.Empty, error) {
	err := s.trackService.DeleteById(ctx, models.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
