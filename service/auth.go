package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"

	"golang.org/x/crypto/bcrypt"
)

// AuthServiceInterface defines the interface for auth service
type AuthServiceInterface interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*model.User, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

// AuthService implements AuthServiceInterface
type AuthService struct {
	userRepo repository.UserRepositoryInterface
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Login authenticates a user
func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*model.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

// Register creates a new user
func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*model.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Check if email already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	user := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
		Role:     "admin",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepo.GetByID(ctx, id)
}
