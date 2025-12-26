package repository

import (
	"context"
	"session-19/model"

	"github.com/stretchr/testify/mock"
)

// MockPortfolioRepository is a mock implementation of PortfolioRepositoryInterface using testify/mock
type MockPortfolioRepository struct {
	mock.Mock
}

// Profile operations
func (m *MockPortfolioRepository) GetProfile(ctx context.Context) (*model.Profile, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Profile), args.Error(1)
}

func (m *MockPortfolioRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	args := m.Called(ctx, profile)
	return args.Error(0)
}

func (m *MockPortfolioRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	args := m.Called(ctx, profile)
	return args.Error(0)
}

// Experience operations
func (m *MockPortfolioRepository) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Experience), args.Error(1)
}

func (m *MockPortfolioRepository) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Experience), args.Error(1)
}

func (m *MockPortfolioRepository) CreateExperience(ctx context.Context, exp *model.Experience) error {
	args := m.Called(ctx, exp)
	return args.Error(0)
}

func (m *MockPortfolioRepository) UpdateExperience(ctx context.Context, exp *model.Experience) error {
	args := m.Called(ctx, exp)
	return args.Error(0)
}

func (m *MockPortfolioRepository) DeleteExperience(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Skill operations
func (m *MockPortfolioRepository) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Skill), args.Error(1)
}

func (m *MockPortfolioRepository) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	args := m.Called(ctx, category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Skill), args.Error(1)
}

func (m *MockPortfolioRepository) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Skill), args.Error(1)
}

func (m *MockPortfolioRepository) CreateSkill(ctx context.Context, skill *model.Skill) error {
	args := m.Called(ctx, skill)
	return args.Error(0)
}

func (m *MockPortfolioRepository) UpdateSkill(ctx context.Context, skill *model.Skill) error {
	args := m.Called(ctx, skill)
	return args.Error(0)
}

func (m *MockPortfolioRepository) DeleteSkill(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Project operations
func (m *MockPortfolioRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *MockPortfolioRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Project), args.Error(1)
}

func (m *MockPortfolioRepository) CreateProject(ctx context.Context, project *model.Project) error {
	args := m.Called(ctx, project)
	return args.Error(0)
}

func (m *MockPortfolioRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	args := m.Called(ctx, project)
	return args.Error(0)
}

func (m *MockPortfolioRepository) DeleteProject(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Publication operations
func (m *MockPortfolioRepository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Publication), args.Error(1)
}

func (m *MockPortfolioRepository) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Publication), args.Error(1)
}

func (m *MockPortfolioRepository) CreatePublication(ctx context.Context, pub *model.Publication) error {
	args := m.Called(ctx, pub)
	return args.Error(0)
}

func (m *MockPortfolioRepository) UpdatePublication(ctx context.Context, pub *model.Publication) error {
	args := m.Called(ctx, pub)
	return args.Error(0)
}

func (m *MockPortfolioRepository) DeletePublication(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Portfolio operations
func (m *MockPortfolioRepository) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.PortfolioData), args.Error(1)
}
