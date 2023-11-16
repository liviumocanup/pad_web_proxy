package adapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"user_service/config"
	"user_service/models"
	"user_service/proto"
	"user_service/services"
)

func NewGrpcServer(cfg *config.Config, userService services.UserService, logger zerolog.Logger) (*grpc.Server, net.Listener, *prometheus.Registry, error) {
	log.Info().Msg("Creating new gRPC server for user service")

	// Setup metrics.
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	server := &grpcServer{
		userService: userService,
		semaphore:   make(chan struct{}, cfg.ConcurrentLimit),
	}

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		return nil, nil, nil, err
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)),
			srvMetrics.UnaryServerInterceptor(),
		),
	)
	proto.RegisterUserServiceServer(srv, server)

	return srv, listener, reg, nil
}

type grpcServer struct {
	proto.UnsafeUserServiceServer
	userService services.UserService
	semaphore   chan struct{}
}

func (s *grpcServer) Register(ctx context.Context, request *proto.UserRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	internalReq := models.UserRequest{
		Username: request.Username,
		Password: request.Password,
	}

	err := s.userService.Register(ctx, internalReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) Login(ctx context.Context, request *proto.UserRequest) (*proto.JWT, error) {
	s.acquire()
	defer s.release()

	internalReq := models.UserRequest{
		Username: request.Username,
		Password: request.Password,
	}

	token, err := s.userService.Login(ctx, internalReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	protoToken := proto.JWT{
		Token: token.Token,
	}

	return &protoToken, nil
}

func (s *grpcServer) Validate(ctx context.Context, jwt *proto.JWT) (*proto.UserResponse, error) {
	s.acquire()
	defer s.release()

	internalReq := models.JWT{
		Token: jwt.Token,
	}

	user, err := s.userService.Validate(ctx, internalReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) FindById(ctx context.Context, request *proto.UserIdRequest) (*proto.UserResponse, error) {
	s.acquire()
	defer s.release()

	internalReq := models.UserIdRequest{
		Id: request.Id,
	}

	user, err := s.userService.FindById(ctx, internalReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) FindByUsername(ctx context.Context, request *proto.UsernameRequest) (*proto.UserResponse, error) {
	s.acquire()
	defer s.release()

	internalReq := models.UsernameRequest{
		Username: request.Username,
	}

	user, err := s.userService.FindByUsername(ctx, internalReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	protoUser := proto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}

	return &protoUser, nil
}

func (s *grpcServer) DeleteById(ctx context.Context, request *proto.UserIdRequest) (*empty.Empty, error) {
	s.acquire()
	defer s.release()

	err := s.userService.DeleteById(ctx, models.UserIdRequest{
		Id: request.Id,
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *grpcServer) FindAll(ctx context.Context, _ *empty.Empty) (*proto.UserListResponse, error) {
	s.acquire()
	defer s.release()

	users, err := s.userService.FindAll(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
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

func (s *grpcServer) Status(ctx context.Context, _ *empty.Empty) (*proto.StatusResponse, error) {
	s.acquire()
	defer s.release()

	serviceStatus, err := s.userService.Status(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		}
		return nil, err
	}

	return &proto.StatusResponse{Status: serviceStatus}, nil
}

func (s *grpcServer) acquire() {
	s.semaphore <- struct{}{}
}

func (s *grpcServer) release() {
	<-s.semaphore
}

func InterceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l := l.With().Fields(fields).Logger()

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Msg(msg)
		case logging.LevelInfo:
			l.Info().Msg(msg)
		case logging.LevelWarn:
			l.Warn().Msg(msg)
		case logging.LevelError:
			l.Error().Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
