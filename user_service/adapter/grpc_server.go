package adapter

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_service/config"
	"user_service/models"
	"user_service/proto"
	"user_service/services"
)

func NewGrpcServer(cfg *config.Config, userService services.UserService) (*grpc.Server, net.Listener, error) {
	log.Println("Creating new gRPC server for user service")

	server := &grpcServer{
		userService: userService,
	}

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, nil, err
	}

	srv := grpc.NewServer()
	proto.RegisterUserServiceServer(srv, server)

	return srv, listener, nil
}

type grpcServer struct {
	proto.UnsafeUserServiceServer
	userService services.UserService
}

func (s *grpcServer) Register(ctx context.Context, request *proto.UserRequest) (*empty.Empty, error) {
	internalReq := models.UserRequest{
		Username: request.Username,
		Password: request.Password,
	}

	err := s.userService.Register(ctx, internalReq)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) Login(ctx context.Context, request *proto.UserRequest) (*proto.JWT, error) {
	internalReq := models.UserRequest{
		Username: request.Username,
		Password: request.Password,
	}

	token, err := s.userService.Login(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	protoToken := proto.JWT{
		Token: token.Token,
	}

	return &protoToken, nil
}

func (s *grpcServer) Validate(ctx context.Context, jwt *proto.JWT) (*proto.UserResponse, error) {
	internalReq := models.JWT{
		Token: jwt.Token,
	}

	user, err := s.userService.Validate(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) FindById(ctx context.Context, request *proto.IdRequest) (*proto.UserResponse, error) {
	internalReq := models.IdRequest{
		Id: request.Id,
	}

	user, err := s.userService.FindById(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) FindByUsername(ctx context.Context, request *proto.UsernameRequest) (*proto.UserResponse, error) {
	internalReq := models.UsernameRequest{
		Username: request.Username,
	}

	user, err := s.userService.FindByUsername(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) FindAll(ctx context.Context, empty *empty.Empty) (*proto.UserListResponse, error) {
	users, err := s.userService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var protoUsers []*proto.UserResponse
	for _, user := range users {
		protoUser := &proto.UserResponse{
			Id:       user.Id,
			Username: user.Username,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	return &proto.UserListResponse{
		Users: protoUsers,
	}, nil
}
