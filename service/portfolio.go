package service

import (
	"context"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
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

// PortfolioService implements PortfolioServiceInterface by aggregating all services
type PortfolioService struct {
	profileSvc     ProfileServiceInterface
	experienceSvc  ExperienceServiceInterface
	skillSvc       SkillServiceInterface
	projectSvc     ProjectServiceInterface
	publicationSvc PublicationServiceInterface
	contactSvc     ContactServiceInterface
	repo           repository.PortfolioRepositoryInterface
}

// NewPortfolioService creates a new portfolio service
func NewPortfolioService(repo repository.PortfolioRepositoryInterface) PortfolioServiceInterface {
	return &PortfolioService{
		profileSvc:     NewProfileService(repo),
		experienceSvc:  NewExperienceService(repo),
		skillSvc:       NewSkillService(repo),
		projectSvc:     NewProjectService(repo),
		publicationSvc: NewPublicationService(repo),
		contactSvc:     NewContactService(),
		repo:           repo,
	}
}

// Profile operations
func (s *PortfolioService) GetProfile(ctx context.Context) (*model.Profile, error) {
	return s.profileSvc.GetProfile(ctx)
}

func (s *PortfolioService) CreateProfile(ctx context.Context, req *dto.ProfileRequest) (*model.Profile, error) {
	return s.profileSvc.CreateProfile(ctx, req)
}

func (s *PortfolioService) UpdateProfile(ctx context.Context, id int64, req *dto.ProfileRequest) (*model.Profile, error) {
	return s.profileSvc.UpdateProfile(ctx, id, req)
}

// Experience operations
func (s *PortfolioService) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	return s.experienceSvc.GetAllExperiences(ctx)
}

func (s *PortfolioService) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	return s.experienceSvc.GetExperienceByID(ctx, id)
}

func (s *PortfolioService) CreateExperience(ctx context.Context, req *dto.ExperienceRequest) (*model.Experience, error) {
	return s.experienceSvc.CreateExperience(ctx, req)
}

func (s *PortfolioService) UpdateExperience(ctx context.Context, id int64, req *dto.ExperienceRequest) (*model.Experience, error) {
	return s.experienceSvc.UpdateExperience(ctx, id, req)
}

func (s *PortfolioService) DeleteExperience(ctx context.Context, id int64) error {
	return s.experienceSvc.DeleteExperience(ctx, id)
}

// Skill operations
func (s *PortfolioService) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	return s.skillSvc.GetAllSkills(ctx)
}

func (s *PortfolioService) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	return s.skillSvc.GetSkillsByCategory(ctx, category)
}

func (s *PortfolioService) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	return s.skillSvc.GetSkillByID(ctx, id)
}

func (s *PortfolioService) CreateSkill(ctx context.Context, req *dto.SkillRequest) (*model.Skill, error) {
	return s.skillSvc.CreateSkill(ctx, req)
}

func (s *PortfolioService) UpdateSkill(ctx context.Context, id int64, req *dto.SkillRequest) (*model.Skill, error) {
	return s.skillSvc.UpdateSkill(ctx, id, req)
}

func (s *PortfolioService) DeleteSkill(ctx context.Context, id int64) error {
	return s.skillSvc.DeleteSkill(ctx, id)
}

// Project operations
func (s *PortfolioService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.projectSvc.GetAllProjects(ctx)
}

func (s *PortfolioService) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	return s.projectSvc.GetProjectByID(ctx, id)
}

func (s *PortfolioService) CreateProject(ctx context.Context, req *dto.ProjectRequest) (*model.Project, error) {
	return s.projectSvc.CreateProject(ctx, req)
}

func (s *PortfolioService) UpdateProject(ctx context.Context, id int64, req *dto.ProjectRequest) (*model.Project, error) {
	return s.projectSvc.UpdateProject(ctx, id, req)
}

func (s *PortfolioService) DeleteProject(ctx context.Context, id int64) error {
	return s.projectSvc.DeleteProject(ctx, id)
}

// Publication operations
func (s *PortfolioService) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return s.publicationSvc.GetAllPublications(ctx)
}

func (s *PortfolioService) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	return s.publicationSvc.GetPublicationByID(ctx, id)
}

func (s *PortfolioService) CreatePublication(ctx context.Context, req *dto.PublicationRequest) (*model.Publication, error) {
	return s.publicationSvc.CreatePublication(ctx, req)
}

func (s *PortfolioService) UpdatePublication(ctx context.Context, id int64, req *dto.PublicationRequest) (*model.Publication, error) {
	return s.publicationSvc.UpdatePublication(ctx, id, req)
}

func (s *PortfolioService) DeletePublication(ctx context.Context, id int64) error {
	return s.publicationSvc.DeletePublication(ctx, id)
}

// Full portfolio data
func (s *PortfolioService) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	return s.repo.GetPortfolioData(ctx)
}

// Contact
func (s *PortfolioService) SubmitContact(ctx context.Context, req *dto.ContactRequest) error {
	return s.contactSvc.SubmitContact(ctx, req)
}
