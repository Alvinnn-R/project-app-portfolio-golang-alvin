package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"strings"
)

// SkillServiceInterface defines the interface for skill service
type SkillServiceInterface interface {
	GetAllSkills(ctx context.Context) ([]model.Skill, error)
	GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*model.Skill, error)
	CreateSkill(ctx context.Context, req *dto.SkillRequest) (*model.Skill, error)
	UpdateSkill(ctx context.Context, id int64, req *dto.SkillRequest) (*model.Skill, error)
	DeleteSkill(ctx context.Context, id int64) error
}

// SkillService implements SkillServiceInterface
type SkillService struct {
	repo repository.SkillRepositoryInterface
}

// NewSkillService creates a new skill service
func NewSkillService(repo repository.SkillRepositoryInterface) SkillServiceInterface {
	return &SkillService{
		repo: repo,
	}
}

// GetAllSkills retrieves all skills
func (s *SkillService) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	return s.repo.GetAllSkills(ctx)
}

// GetSkillsByCategory retrieves skills by category
func (s *SkillService) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	if category == "" {
		return nil, errors.New("category is required")
	}
	return s.repo.GetSkillsByCategory(ctx, category)
}

// GetSkillByID retrieves a skill by ID
func (s *SkillService) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	if id <= 0 {
		return nil, errors.New("invalid skill ID")
	}
	return s.repo.GetSkillByID(ctx, id)
}

// CreateSkill creates a new skill
func (s *SkillService) CreateSkill(ctx context.Context, req *dto.SkillRequest) (*model.Skill, error) {
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

// UpdateSkill updates a skill
func (s *SkillService) UpdateSkill(ctx context.Context, id int64, req *dto.SkillRequest) (*model.Skill, error) {
	if id <= 0 {
		return nil, errors.New("invalid skill ID")
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
func (s *SkillService) DeleteSkill(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid skill ID")
	}
	return s.repo.DeleteSkill(ctx, id)
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
