package service

import (
	"context"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// ProfileServiceInterface defines the interface for profile service
type ProfileServiceInterface interface {
	GetProfile(ctx context.Context) (*model.Profile, error)
	CreateProfile(ctx context.Context, req *dto.ProfileRequest) (*model.Profile, error)
	UpdateProfile(ctx context.Context, id int64, req *dto.ProfileRequest) (*model.Profile, error)
}

// ProfileService implements ProfileServiceInterface
type ProfileService struct {
	repo repository.PortfolioRepositoryInterface
}

// NewProfileService creates a new profile service
func NewProfileService(repo repository.PortfolioRepositoryInterface) ProfileServiceInterface {
	return &ProfileService{
		repo: repo,
	}
}

// GetProfile retrieves the main profile
func (s *ProfileService) GetProfile(ctx context.Context) (*model.Profile, error) {
	return s.repo.GetProfile(ctx)
}

// CreateProfile creates a new profile
func (s *ProfileService) CreateProfile(ctx context.Context, req *dto.ProfileRequest) (*model.Profile, error) {
	profile := &model.Profile{
		Name:        strings.TrimSpace(req.Name),
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		PhotoURL:    strings.TrimSpace(req.PhotoURL),
		Email:       strings.TrimSpace(req.Email),
		LinkedInURL: strings.TrimSpace(req.LinkedInURL),
		GithubURL:   strings.TrimSpace(req.GithubURL),
		CVURL:       strings.TrimSpace(req.CVURL),
	}

	if err := s.repo.CreateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}

// UpdateProfile updates the profile
func (s *ProfileService) UpdateProfile(ctx context.Context, id int64, req *dto.ProfileRequest) (*model.Profile, error) {
	profile := &model.Profile{
		ID:          id,
		Name:        strings.TrimSpace(req.Name),
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		PhotoURL:    strings.TrimSpace(req.PhotoURL),
		Email:       strings.TrimSpace(req.Email),
		LinkedInURL: strings.TrimSpace(req.LinkedInURL),
		GithubURL:   strings.TrimSpace(req.GithubURL),
		CVURL:       strings.TrimSpace(req.CVURL),
	}

	if err := s.repo.UpdateProfile(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}
