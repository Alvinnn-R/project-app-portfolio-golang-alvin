package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// PortfolioRepositoryInterface defines the interface for portfolio repository
type PortfolioRepositoryInterface interface {
	// Profile operations
	GetProfile(ctx context.Context) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error

	// Experience operations
	GetAllExperiences(ctx context.Context) ([]model.Experience, error)
	GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error)
	CreateExperience(ctx context.Context, exp *model.Experience) error
	UpdateExperience(ctx context.Context, exp *model.Experience) error
	DeleteExperience(ctx context.Context, id int64) error

	// Skill operations
	GetAllSkills(ctx context.Context) ([]model.Skill, error)
	GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*model.Skill, error)
	CreateSkill(ctx context.Context, skill *model.Skill) error
	UpdateSkill(ctx context.Context, skill *model.Skill) error
	DeleteSkill(ctx context.Context, id int64) error

	// Project operations
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) error
	UpdateProject(ctx context.Context, project *model.Project) error
	DeleteProject(ctx context.Context, id int64) error

	// Publication operations
	GetAllPublications(ctx context.Context) ([]model.Publication, error)
	GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error)
	CreatePublication(ctx context.Context, pub *model.Publication) error
	UpdatePublication(ctx context.Context, pub *model.Publication) error
	DeletePublication(ctx context.Context, id int64) error

	// Full portfolio data
	GetPortfolioData(ctx context.Context) (*model.PortfolioData, error)
}

// PortfolioRepository implements PortfolioRepositoryInterface by aggregating other repositories
type PortfolioRepository struct {
	profileRepo     ProfileRepositoryInterface
	experienceRepo  ExperienceRepositoryInterface
	skillRepo       SkillRepositoryInterface
	projectRepo     ProjectRepositoryInterface
	publicationRepo PublicationRepositoryInterface
	log             *zap.Logger
}

// NewPortfolioRepository creates a new portfolio repository
func NewPortfolioRepository(db database.PgxIface, log *zap.Logger) PortfolioRepositoryInterface {
	return &PortfolioRepository{
		profileRepo:     NewProfileRepository(db, log),
		experienceRepo:  NewExperienceRepository(db, log),
		skillRepo:       NewSkillRepository(db, log),
		projectRepo:     NewProjectRepository(db, log),
		publicationRepo: NewPublicationRepository(db, log),
		log:             log,
	}
}

// GetProfile retrieves the main profile
func (r *PortfolioRepository) GetProfile(ctx context.Context) (*model.Profile, error) {
	return r.profileRepo.GetProfile(ctx)
}

// CreateProfile creates a new profile
func (r *PortfolioRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	return r.profileRepo.CreateProfile(ctx, profile)
}

// UpdateProfile updates the profile
func (r *PortfolioRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	return r.profileRepo.UpdateProfile(ctx, profile)
}

// GetAllExperiences retrieves all experiences
func (r *PortfolioRepository) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	return r.experienceRepo.GetAllExperiences(ctx)
}

// GetExperienceByID retrieves an experience by ID
func (r *PortfolioRepository) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	return r.experienceRepo.GetExperienceByID(ctx, id)
}

// CreateExperience creates a new experience
func (r *PortfolioRepository) CreateExperience(ctx context.Context, exp *model.Experience) error {
	return r.experienceRepo.CreateExperience(ctx, exp)
}

// UpdateExperience updates an experience
func (r *PortfolioRepository) UpdateExperience(ctx context.Context, exp *model.Experience) error {
	return r.experienceRepo.UpdateExperience(ctx, exp)
}

// DeleteExperience deletes an experience
func (r *PortfolioRepository) DeleteExperience(ctx context.Context, id int64) error {
	return r.experienceRepo.DeleteExperience(ctx, id)
}

// GetAllSkills retrieves all skills
func (r *PortfolioRepository) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	return r.skillRepo.GetAllSkills(ctx)
}

// GetSkillsByCategory retrieves skills by category
func (r *PortfolioRepository) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	return r.skillRepo.GetSkillsByCategory(ctx, category)
}

// GetSkillByID retrieves a skill by ID
func (r *PortfolioRepository) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	return r.skillRepo.GetSkillByID(ctx, id)
}

// CreateSkill creates a new skill
func (r *PortfolioRepository) CreateSkill(ctx context.Context, skill *model.Skill) error {
	return r.skillRepo.CreateSkill(ctx, skill)
}

// UpdateSkill updates a skill
func (r *PortfolioRepository) UpdateSkill(ctx context.Context, skill *model.Skill) error {
	return r.skillRepo.UpdateSkill(ctx, skill)
}

// DeleteSkill deletes a skill
func (r *PortfolioRepository) DeleteSkill(ctx context.Context, id int64) error {
	return r.skillRepo.DeleteSkill(ctx, id)
}

// GetAllProjects retrieves all projects
func (r *PortfolioRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return r.projectRepo.GetAllProjects(ctx)
}

// GetProjectByID retrieves a project by ID
func (r *PortfolioRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	return r.projectRepo.GetProjectByID(ctx, id)
}

// CreateProject creates a new project
func (r *PortfolioRepository) CreateProject(ctx context.Context, project *model.Project) error {
	return r.projectRepo.CreateProject(ctx, project)
}

// UpdateProject updates a project
func (r *PortfolioRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	return r.projectRepo.UpdateProject(ctx, project)
}

// DeleteProject deletes a project
func (r *PortfolioRepository) DeleteProject(ctx context.Context, id int64) error {
	return r.projectRepo.DeleteProject(ctx, id)
}

// GetAllPublications retrieves all publications
func (r *PortfolioRepository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return r.publicationRepo.GetAllPublications(ctx)
}

// GetPublicationByID retrieves a publication by ID
func (r *PortfolioRepository) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	return r.publicationRepo.GetPublicationByID(ctx, id)
}

// CreatePublication creates a new publication
func (r *PortfolioRepository) CreatePublication(ctx context.Context, pub *model.Publication) error {
	return r.publicationRepo.CreatePublication(ctx, pub)
}

// UpdatePublication updates a publication
func (r *PortfolioRepository) UpdatePublication(ctx context.Context, pub *model.Publication) error {
	return r.publicationRepo.UpdatePublication(ctx, pub)
}

// DeletePublication deletes a publication
func (r *PortfolioRepository) DeletePublication(ctx context.Context, id int64) error {
	return r.publicationRepo.DeletePublication(ctx, id)
}

// GetPortfolioData retrieves all portfolio data in one call
func (r *PortfolioRepository) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	data := &model.PortfolioData{}

	// Get profile
	profile, err := r.GetProfile(ctx)
	if err != nil {
		r.log.Warn("No profile found, using empty profile", zap.Error(err))
		data.Profile = model.Profile{}
	} else {
		data.Profile = *profile
	}

	// Get experiences
	experiences, err := r.GetAllExperiences(ctx)
	if err != nil {
		r.log.Warn("Failed to get experiences", zap.Error(err))
		data.Experiences = []model.Experience{}
	} else {
		data.Experiences = experiences
	}

	// Get skills and group by category
	skills, err := r.GetAllSkills(ctx)
	if err != nil {
		r.log.Warn("Failed to get skills", zap.Error(err))
		data.Skills = make(map[string][]model.Skill)
	} else {
		skillMap := make(map[string][]model.Skill)
		for _, skill := range skills {
			skillMap[skill.Category] = append(skillMap[skill.Category], skill)
		}
		data.Skills = skillMap
	}

	// Get projects
	projects, err := r.GetAllProjects(ctx)
	if err != nil {
		r.log.Warn("Failed to get projects", zap.Error(err))
		data.Projects = []model.Project{}
	} else {
		data.Projects = projects
	}

	// Get publications
	publications, err := r.GetAllPublications(ctx)
	if err != nil {
		r.log.Warn("Failed to get publications", zap.Error(err))
		data.Publications = []model.Publication{}
	} else {
		data.Publications = publications
	}

	return data, nil
}
