package services

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strconv"
	"user_service/auth"
	"user_service/config"
	"user_service/models"
	"user_service/repositories"
)

type UserService interface {
	Register(ctx context.Context, request models.UserRequest) error
	Login(ctx context.Context, request models.UserRequest) (*models.JWT, error)
	Validate(ctx context.Context, jwt models.JWT) (*models.UserResponse, error)
	FindById(ctx context.Context, request models.UserIdRequest) (*models.UserResponse, error)
	FindByUsername(ctx context.Context, request models.UsernameRequest) (*models.UserResponse, error)
	DeleteById(ctx context.Context, request models.UserIdRequest) error
	FindAll(ctx context.Context) ([]*models.UserResponse, error)
	Status(ctx context.Context) (string, error)
}

type UserServiceServer struct {
	repository *repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository, cfg *config.Config) UserService {
	log.Info().Msg("Creating user service")

	return &userService{
		repository: repository,
		cfg:        cfg,
	}
}

type userService struct {
	repository repositories.UserRepository
	cfg        *config.Config
}

func (s *userService) Register(ctx context.Context, request models.UserRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	existingUser, err := s.repository.FindByUsername(request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Error(codes.Internal, "internal error")
	}
	if existingUser != nil {
		return status.Error(codes.AlreadyExists, "user already exists")
	}

	hashedPassword := auth.HashPassword(request.Password)

	user := &models.User{
		Username: request.Username,
		Password: hashedPassword,
	}

	if err := s.repository.Create(user); err != nil {
		return status.Error(codes.Internal, "failed to create user")
	}

	return nil
}

func (s *userService) Login(ctx context.Context, request models.UserRequest) (*models.JWT, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	user, err := s.repository.FindByUsername(request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	if !auth.CheckPassword(user.Password, request.Password) {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, []byte(s.cfg.JWTSecret), s.cfg.JWTDuration)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &models.JWT{Token: token}, nil
}

func (s *userService) Validate(ctx context.Context, jwt models.JWT) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	// TEST FOR TIMEOUT
	//select {
	//case <-time.After(10 * time.Second):
	//	log.Info().Msg("Sleep Over.....")
	//case <-ctx.Done():
	//	return nil, ctx.Err()
	//}

	userId, err := auth.ValidateToken(jwt.Token, []byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	user, err := s.repository.FindByID(strconv.Itoa(int(userId)))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &models.UserResponse{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	}, nil
}

func (s *userService) FindById(ctx context.Context, request models.UserIdRequest) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	// TEST FOR CONCURRENT TASK LIMIT
	//select {
	//case <-time.After(1 * time.Second):
	//	log.Info().Msg("Sleep Over.....")
	//case <-ctx.Done():
	//	return nil, ctx.Err()
	//}

	user, err := s.repository.FindByID(request.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &models.UserResponse{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	}, nil
}

func (s *userService) FindByUsername(ctx context.Context, request models.UsernameRequest) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	user, err := s.repository.FindByUsername(request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &models.UserResponse{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	}, nil
}

func (s *userService) DeleteById(ctx context.Context, request models.UserIdRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	if err := s.repository.Delete(request.Id); err != nil {
		return status.Error(codes.Internal, "failed to delete user")
	}
	return nil
}

func (s *userService) FindAll(ctx context.Context) ([]*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	users, err := s.repository.FindAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	var userResponses []*models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &models.UserResponse{
			Id:       strconv.Itoa(int(user.ID)),
			Username: user.Username,
		})
	}

	return userResponses, nil
}

func (s *userService) Status(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.RequestTimeout)
	defer cancel()

	return "ok", nil
}

func (s *userService) mustEmbedUnimplementedUserServiceServer() {}
