package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// PortfolioServiceInterface defines the interface for portfolio service
type PortfolioServiceInterface interface {
	// Profile operations
	GetProfile(ctx context.Context) (*model.Profile, error)
	CreateProfile(ctx context.Context, req *dto.ProfileRequest) (*model.Profile, error)
	UpdateProfile(ctx context.Context, id int64, req *dto.ProfileRequest) (*model.Profile, error)

	// Experience operations
	GetAllExperiences(ctx context.Context) ([]model.Experience, error)
	GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error)
	CreateExperience(ctx context.Context, req *dto.ExperienceRequest) (*model.Experience, error)
	UpdateExperience(ctx context.Context, id int64, req *dto.ExperienceRequest) (*model.Experience, error)
	DeleteExperience(ctx context.Context, id int64) error

	// Skill operations
	GetAllSkills(ctx context.Context) ([]model.Skill, error)
	GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*model.Skill, error)
	CreateSkill(ctx context.Context, req *dto.SkillRequest) (*model.Skill, error)
	UpdateSkill(ctx context.Context, id int64, req *dto.SkillRequest) (*model.Skill, error)
	DeleteSkill(ctx context.Context, id int64) error

	// Project operations
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, req *dto.ProjectRequest) (*model.Project, error)
	UpdateProject(ctx context.Context, id int64, req *dto.ProjectRequest) (*model.Project, error)
	DeleteProject(ctx context.Context, id int64) error

	// Publication operations
	GetAllPublications(ctx context.Context) ([]model.Publication, error)
	GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error)
	CreatePublication(ctx context.Context, req *dto.PublicationRequest) (*model.Publication, error)
	UpdatePublication(ctx context.Context, id int64, req *dto.PublicationRequest) (*model.Publication, error)
	DeletePublication(ctx context.Context, id int64) error

	// Full portfolio data
	GetPortfolioData(ctx context.Context) (*model.PortfolioData, error)

	// Contact
	SubmitContact(ctx context.Context, req *dto.ContactRequest) error
}

// PortfolioService implements PortfolioServiceInterface
type PortfolioService struct {
	repo repository.PortfolioRepositoryInterface
}

// NewPortfolioService creates a new portfolio service
func NewPortfolioService(repo repository.PortfolioRepositoryInterface) PortfolioServiceInterface {
	return &PortfolioService{
		repo: repo,
	}
}

// GetProfile retrieves the main profile
func (s *PortfolioService) GetProfile(ctx context.Context) (*model.Profile, error) {
	return s.repo.GetProfile(ctx)
}

// CreateProfile creates a new profile with validation
func (s *PortfolioService) CreateProfile(ctx context.Context, req *dto.ProfileRequest) (*model.Profile, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

// UpdateProfile updates the profile with validation
func (s *PortfolioService) UpdateProfile(ctx context.Context, id int64, req *dto.ProfileRequest) (*model.Profile, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

// GetAllExperiences retrieves all experiences
func (s *PortfolioService) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	return s.repo.GetAllExperiences(ctx)
}

// GetExperienceByID retrieves an experience by ID
func (s *PortfolioService) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	if id <= 0 {
		return nil, errors.New("invalid experience ID")
	}
	return s.repo.GetExperienceByID(ctx, id)
}

// CreateExperience creates a new experience with validation
func (s *PortfolioService) CreateExperience(ctx context.Context, req *dto.ExperienceRequest) (*model.Experience, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

// UpdateExperience updates an experience with validation
func (s *PortfolioService) UpdateExperience(ctx context.Context, id int64, req *dto.ExperienceRequest) (*model.Experience, error) {
	if id <= 0 {
		return nil, errors.New("invalid experience ID")
	}

	if err := req.Validate(); err != nil {
		return nil, err
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
func (s *PortfolioService) DeleteExperience(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid experience ID")
	}
	return s.repo.DeleteExperience(ctx, id)
}

// GetAllSkills retrieves all skills
func (s *PortfolioService) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	return s.repo.GetAllSkills(ctx)
}

// GetSkillsByCategory retrieves skills by category
func (s *PortfolioService) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	if category == "" {
		return nil, errors.New("category is required")
	}
	return s.repo.GetSkillsByCategory(ctx, category)
}

// GetSkillByID retrieves a skill by ID
func (s *PortfolioService) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	if id <= 0 {
		return nil, errors.New("invalid skill ID")
	}
	return s.repo.GetSkillByID(ctx, id)
}

// CreateSkill creates a new skill with validation
func (s *PortfolioService) CreateSkill(ctx context.Context, req *dto.SkillRequest) (*model.Skill, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	skill := &model.Skill{
		Category: strings.TrimSpace(req.Category),
		Name:     strings.TrimSpace(req.Name),
		Level:    strings.TrimSpace(req.Level),
		Color:    getColorForLevel(req.Level, req.Color),
	}

	if err := s.repo.CreateSkill(ctx, skill); err != nil {
		return nil, err
	}

	return skill, nil
}

// UpdateSkill updates a skill with validation
func (s *PortfolioService) UpdateSkill(ctx context.Context, id int64, req *dto.SkillRequest) (*model.Skill, error) {
	if id <= 0 {
		return nil, errors.New("invalid skill ID")
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	skill := &model.Skill{
		ID:       id,
		Category: strings.TrimSpace(req.Category),
		Name:     strings.TrimSpace(req.Name),
		Level:    strings.TrimSpace(req.Level),
		Color:    getColorForLevel(req.Level, req.Color),
	}

	if err := s.repo.UpdateSkill(ctx, skill); err != nil {
		return nil, err
	}

	return skill, nil
}

// DeleteSkill deletes a skill
func (s *PortfolioService) DeleteSkill(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid skill ID")
	}
	return s.repo.DeleteSkill(ctx, id)
}

// GetAllProjects retrieves all projects
func (s *PortfolioService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.repo.GetAllProjects(ctx)
}

// GetProjectByID retrieves a project by ID
func (s *PortfolioService) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	if id <= 0 {
		return nil, errors.New("invalid project ID")
	}
	return s.repo.GetProjectByID(ctx, id)
}

// CreateProject creates a new project with validation
func (s *PortfolioService) CreateProject(ctx context.Context, req *dto.ProjectRequest) (*model.Project, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

// UpdateProject updates a project with validation
func (s *PortfolioService) UpdateProject(ctx context.Context, id int64, req *dto.ProjectRequest) (*model.Project, error) {
	if id <= 0 {
		return nil, errors.New("invalid project ID")
	}

	if err := req.Validate(); err != nil {
		return nil, err
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
func (s *PortfolioService) DeleteProject(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid project ID")
	}
	return s.repo.DeleteProject(ctx, id)
}

// GetAllPublications retrieves all publications
func (s *PortfolioService) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return s.repo.GetAllPublications(ctx)
}

// GetPublicationByID retrieves a publication by ID
func (s *PortfolioService) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	if id <= 0 {
		return nil, errors.New("invalid publication ID")
	}
	return s.repo.GetPublicationByID(ctx, id)
}

// CreatePublication creates a new publication with validation
func (s *PortfolioService) CreatePublication(ctx context.Context, req *dto.PublicationRequest) (*model.Publication, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	pub := &model.Publication{
		Title:          strings.TrimSpace(req.Title),
		Authors:        strings.TrimSpace(req.Authors),
		Journal:        strings.TrimSpace(req.Journal),
		Year:           req.Year,
		Description:    strings.TrimSpace(req.Description),
		ImageURL:       strings.TrimSpace(req.ImageURL),
		PublicationURL: strings.TrimSpace(req.PublicationURL),
		Color:          getDefaultColor(req.Color, "red"),
	}

	if err := s.repo.CreatePublication(ctx, pub); err != nil {
		return nil, err
	}

	return pub, nil
}

// UpdatePublication updates a publication with validation
func (s *PortfolioService) UpdatePublication(ctx context.Context, id int64, req *dto.PublicationRequest) (*model.Publication, error) {
	if id <= 0 {
		return nil, errors.New("invalid publication ID")
	}

	if err := req.Validate(); err != nil {
		return nil, err
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
		Color:          getDefaultColor(req.Color, "red"),
	}

	if err := s.repo.UpdatePublication(ctx, pub); err != nil {
		return nil, err
	}

	return pub, nil
}

// DeletePublication deletes a publication
func (s *PortfolioService) DeletePublication(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid publication ID")
	}
	return s.repo.DeletePublication(ctx, id)
}

// GetPortfolioData retrieves all portfolio data
func (s *PortfolioService) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	return s.repo.GetPortfolioData(ctx)
}

// SubmitContact handles contact form submission
func (s *PortfolioService) SubmitContact(ctx context.Context, req *dto.ContactRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	// In a real application, you would send an email or store the contact message
	// For now, we just validate the request
	return nil
}

// Helper functions

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

// getColorForLevel returns a color based on skill level
func getColorForLevel(level, defaultColor string) string {
	if defaultColor != "" {
		return defaultColor
	}
	switch level {
	case "advanced":
		return "black"
	case "intermediate":
		return "gray"
	case "beginner":
		return "white"
	default:
		return "gray"
	}
}

// getDefaultColor returns the provided color or default if empty
func getDefaultColor(color, defaultColor string) string {
	if color != "" {
		return color
	}
	return defaultColor
}
