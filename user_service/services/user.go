package services

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
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
	FindById(ctx context.Context, request models.IdRequest) (*models.UserResponse, error)
	FindByUsername(ctx context.Context, request models.UsernameRequest) (*models.UserResponse, error)
	FindAll(ctx context.Context) ([]*models.UserResponse, error)
}

type UserServiceServer struct {
	repository *repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository, cfg *config.Config) UserService {
	log.Println("Creating user service")

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

	token, err := auth.GenerateToken(user, []byte(s.cfg.JWTSecret), s.cfg.JWTDuration)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &models.JWT{Token: token}, nil
}

func (s *userService) Validate(ctx context.Context, jwt models.JWT) (*models.UserResponse, error) {
	user, err := auth.ValidateToken(jwt.Token, []byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return &models.UserResponse{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	}, nil
}

func (s *userService) FindById(ctx context.Context, request models.IdRequest) (*models.UserResponse, error) {
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

func (s *userService) FindAll(ctx context.Context) ([]*models.UserResponse, error) {
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

// The following method is required by the newer versions of grpc for forward compatibility.
// You can leave the method body empty.
func (s *userService) mustEmbedUnimplementedUserServiceServer() {}
