package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// ExperienceServiceInterface defines the interface for experience service
type ExperienceServiceInterface interface {
	GetAllExperiences(ctx context.Context) ([]model.Experience, error)
	GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error)
	CreateExperience(ctx context.Context, req *dto.ExperienceRequest) (*model.Experience, error)
	UpdateExperience(ctx context.Context, id int64, req *dto.ExperienceRequest) (*model.Experience, error)
	DeleteExperience(ctx context.Context, id int64) error
}

// ExperienceService implements ExperienceServiceInterface
type ExperienceService struct {
	repo repository.PortfolioRepositoryInterface
}

// NewExperienceService creates a new experience service
func NewExperienceService(repo repository.PortfolioRepositoryInterface) ExperienceServiceInterface {
	return &ExperienceService{
		repo: repo,
	}
}

// GetAllExperiences retrieves all experiences
func (s *ExperienceService) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	return s.repo.GetAllExperiences(ctx)
}

// GetExperienceByID retrieves an experience by ID
func (s *ExperienceService) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	if id <= 0 {
		return nil, errors.New("invalid experience ID")
	}
	return s.repo.GetExperienceByID(ctx, id)
}

// CreateExperience creates a new experience
func (s *ExperienceService) CreateExperience(ctx context.Context, req *dto.ExperienceRequest) (*model.Experience, error) {
	exp := &model.Experience{
		Title:        strings.TrimSpace(req.Title),
		Organization: strings.TrimSpace(req.Organization),
		Period:       strings.TrimSpace(req.Period),
		Description:  strings.TrimSpace(req.Description),
		Type:         strings.TrimSpace(req.Type),
		Color:        getColorForType(req.Type, req.Color),
	}

	if err := s.repo.CreateExperience(ctx, exp); err != nil {
		return nil, err
	}

	return exp, nil
}

// UpdateExperience updates an experience
func (s *ExperienceService) UpdateExperience(ctx context.Context, id int64, req *dto.ExperienceRequest) (*model.Experience, error) {
	if id <= 0 {
		return nil, errors.New("invalid experience ID")
	}

	exp := &model.Experience{
		ID:           id,
		Title:        strings.TrimSpace(req.Title),
		Organization: strings.TrimSpace(req.Organization),
		Period:       strings.TrimSpace(req.Period),
		Description:  strings.TrimSpace(req.Description),
		Type:         strings.TrimSpace(req.Type),
		Color:        getColorForType(req.Type, req.Color),
	}

	if err := s.repo.UpdateExperience(ctx, exp); err != nil {
		return nil, err
	}

	return exp, nil
}

// DeleteExperience deletes an experience
func (s *ExperienceService) DeleteExperience(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid experience ID")
	}
	return s.repo.DeleteExperience(ctx, id)
}

// getColorForType returns a color based on experience type
func getColorForType(expType, defaultColor string) string {
	if defaultColor != "" {
		return defaultColor
	}
	switch expType {
	case "work":
		return "cyan"
	case "internship":
		return "pink"
	case "campus":
		return "yellow"
	case "competition":
		return "purple"
	default:
		return "gray"
	}
}
