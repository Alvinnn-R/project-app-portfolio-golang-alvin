package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// PublicationServiceInterface defines the interface for publication service
type PublicationServiceInterface interface {
	GetAllPublications(ctx context.Context) ([]model.Publication, error)
	GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error)
	CreatePublication(ctx context.Context, req *dto.PublicationRequest) (*model.Publication, error)
	UpdatePublication(ctx context.Context, id int64, req *dto.PublicationRequest) (*model.Publication, error)
	DeletePublication(ctx context.Context, id int64) error
}

// PublicationService implements PublicationServiceInterface
type PublicationService struct {
	repo repository.PublicationRepositoryInterface
}

// NewPublicationService creates a new publication service
func NewPublicationService(repo repository.PublicationRepositoryInterface) PublicationServiceInterface {
	return &PublicationService{
		repo: repo,
	}
}

// GetAllPublications retrieves all publications
func (s *PublicationService) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return s.repo.GetAllPublications(ctx)
}

// GetPublicationByID retrieves a publication by ID
func (s *PublicationService) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	if id <= 0 {
		return nil, errors.New("invalid publication ID")
	}
	return s.repo.GetPublicationByID(ctx, id)
}

// CreatePublication creates a new publication
func (s *PublicationService) CreatePublication(ctx context.Context, req *dto.PublicationRequest) (*model.Publication, error) {
	pub := &model.Publication{
		Title:          strings.TrimSpace(req.Title),
		Authors:        strings.TrimSpace(req.Authors),
		Journal:        strings.TrimSpace(req.Journal),
		Year:           req.Year,
		Description:    strings.TrimSpace(req.Description),
		ImageURL:       strings.TrimSpace(req.ImageURL),
		PublicationURL: strings.TrimSpace(req.PublicationURL),
		Color:          getPublicationDefaultColor(req.Color, "red"),
	}

	if err := s.repo.CreatePublication(ctx, pub); err != nil {
		return nil, err
	}

	return pub, nil
}

// UpdatePublication updates a publication
func (s *PublicationService) UpdatePublication(ctx context.Context, id int64, req *dto.PublicationRequest) (*model.Publication, error) {
	if id <= 0 {
		return nil, errors.New("invalid publication ID")
	}

	pub := &model.Publication{
		ID:             id,
		Title:          strings.TrimSpace(req.Title),
		Authors:        strings.TrimSpace(req.Authors),
		Journal:        strings.TrimSpace(req.Journal),
		Year:           req.Year,
		Description:    strings.TrimSpace(req.Description),
		ImageURL:       strings.TrimSpace(req.ImageURL),
		PublicationURL: strings.TrimSpace(req.PublicationURL),
		Color:          getPublicationDefaultColor(req.Color, "red"),
	}

	if err := s.repo.UpdatePublication(ctx, pub); err != nil {
		return nil, err
	}

	return pub, nil
}

// DeletePublication deletes a publication
func (s *PublicationService) DeletePublication(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid publication ID")
	}
	return s.repo.DeletePublication(ctx, id)
}

// getPublicationDefaultColor returns the provided color or default if empty
func getPublicationDefaultColor(color, defaultColor string) string {
	if color != "" {
		return color
	}
	return defaultColor
}
