package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// ProjectServiceInterface defines the interface for project service
type ProjectServiceInterface interface {
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, req *dto.ProjectRequest) (*model.Project, error)
	UpdateProject(ctx context.Context, id int64, req *dto.ProjectRequest) (*model.Project, error)
	DeleteProject(ctx context.Context, id int64) error
}

// ProjectService implements ProjectServiceInterface
type ProjectService struct {
	repo repository.ProjectRepositoryInterface
}

// NewProjectService creates a new project service
func NewProjectService(repo repository.ProjectRepositoryInterface) ProjectServiceInterface {
	return &ProjectService{
		repo: repo,
	}
}

// GetAllProjects retrieves all projects
func (s *ProjectService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetAllProjects(ctx)
}

// GetProjectByID retrieves a project by ID
func (s *ProjectService) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	if id <= 0 {
		return nil, errors.New("invalid project ID")
	}
	return s.repo.GetProjectByID(ctx, id)
}

// CreateProject creates a new project
func (s *ProjectService) CreateProject(ctx context.Context, req *dto.ProjectRequest) (*model.Project, error) {
	project := &model.Project{
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		ImageURL:    strings.TrimSpace(req.ImageURL),
		ProjectURL:  strings.TrimSpace(req.ProjectURL),
		GithubURL:   strings.TrimSpace(req.GithubURL),
		TechStack:   strings.TrimSpace(req.TechStack),
		Color:       getDefaultColor(req.Color, "cyan"),
		ProfileID:   req.ProfileID,
	}

	if err := s.repo.CreateProject(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

// UpdateProject updates a project
func (s *ProjectService) UpdateProject(ctx context.Context, id int64, req *dto.ProjectRequest) (*model.Project, error) {
	if id <= 0 {
		return nil, errors.New("invalid project ID")
	}

	project := &model.Project{
		ID:          id,
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		ImageURL:    strings.TrimSpace(req.ImageURL),
		ProjectURL:  strings.TrimSpace(req.ProjectURL),
		GithubURL:   strings.TrimSpace(req.GithubURL),
		TechStack:   strings.TrimSpace(req.TechStack),
		Color:       getDefaultColor(req.Color, "cyan"),
		ProfileID:   req.ProfileID,
	}

	if err := s.repo.UpdateProject(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProject deletes a project
func (s *ProjectService) DeleteProject(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid project ID")
	}
	return s.repo.DeleteProject(ctx, id)
}

// getDefaultColor returns the provided color or default if empty
func getDefaultColor(color, defaultColor string) string {
	if color != "" {
		return color
	}
	return defaultColor
}
