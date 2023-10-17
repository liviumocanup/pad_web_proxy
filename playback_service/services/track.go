package services

import (
	"context"
	"google.golang.org/grpc"
	"playback_service/proto"
)

type TrackServiceClient interface {
	GetInfoById(ctx context.Context, trackId string) (*proto.TrackInfoResponse, error)
	Close() error
}

type trackServiceClientImpl struct {
	conn   *grpc.ClientConn
	client proto.TrackServiceClient
}

func NewTrackServiceClient(address string) (TrackServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := proto.NewTrackServiceClient(conn)
	return &trackServiceClientImpl{conn: conn, client: client}, nil
}

func (t *trackServiceClientImpl) GetInfoById(ctx context.Context, trackId string) (*proto.TrackInfoResponse, error) {
	req := &proto.TrackIdRequest{Id: trackId}
	return t.client.GetInfoById(ctx, req)
}

func (t *trackServiceClientImpl) Close() error {
	return t.conn.Close()
}
