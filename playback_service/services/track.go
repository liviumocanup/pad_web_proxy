package services

import (
	"context"
	"google.golang.org/grpc"
	"playback_service/proto"
	"playback_service/utils" // Ensure this is the correct import path
)

type TrackServiceClient interface {
	GetInfoById(ctx context.Context, trackId string) (*proto.TrackInfoResponse, error)
	Close() error
}

type trackServiceClientImpl struct {
	conn           *grpc.ClientConn
	client         proto.TrackServiceClient
	circuitBreaker *utils.CircuitBreaker
}

func NewTrackServiceClient(address string, circuitBreaker *utils.CircuitBreaker) (TrackServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := proto.NewTrackServiceClient(conn)
	return &trackServiceClientImpl{conn: conn, client: client, circuitBreaker: circuitBreaker}, nil
}

func (t *trackServiceClientImpl) GetInfoById(ctx context.Context, trackId string) (*proto.TrackInfoResponse, error) {
	var response *proto.TrackInfoResponse

	err := t.circuitBreaker.Call(func() error {
		req := &proto.TrackIdRequest{Id: trackId}
		var callErr error
		response, callErr = t.client.GetInfoById(ctx, req)
		return callErr
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (t *trackServiceClientImpl) Close() error {
	return t.conn.Close()
}
