package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"testing"
)

// MockPortfolioRepository is a mock implementation of PortfolioRepositoryInterface
type MockPortfolioRepository struct {
	// Profile mock data
	profile    *model.Profile
	profileErr error
	createErr  error
	updateErr  error
	deleteErr  error

	// Experience mock data
	experiences   []model.Experience
	experience    *model.Experience
	experienceErr error

	// Skill mock data
	skills   []model.Skill
	skill    *model.Skill
	skillErr error

	// Project mock data
	projects   []model.Project
	project    *model.Project
	projectErr error

	// Publication mock data
	publications   []model.Publication
	publication    *model.Publication
	publicationErr error

	// Portfolio data
	portfolioData *model.PortfolioData
	portfolioErr  error
}

// Profile operations
func (m *MockPortfolioRepository) GetProfile(ctx context.Context) (*model.Profile, error) {
	return m.profile, m.profileErr
}

func (m *MockPortfolioRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	if m.createErr != nil {
		return m.createErr
	}
	profile.ID = 1
	return nil
}

func (m *MockPortfolioRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	return m.updateErr
}

// Experience operations
func (m *MockPortfolioRepository) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	return m.experiences, m.experienceErr
}

func (m *MockPortfolioRepository) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	return m.experience, m.experienceErr
}

func (m *MockPortfolioRepository) CreateExperience(ctx context.Context, exp *model.Experience) error {
	if m.createErr != nil {
		return m.createErr
	}
	exp.ID = 1
	return nil
}

func (m *MockPortfolioRepository) UpdateExperience(ctx context.Context, exp *model.Experience) error {
	return m.updateErr
}

func (m *MockPortfolioRepository) DeleteExperience(ctx context.Context, id int64) error {
	return m.deleteErr
}

// Skill operations
func (m *MockPortfolioRepository) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	return m.skills, m.skillErr
}

func (m *MockPortfolioRepository) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	return m.skills, m.skillErr
}

func (m *MockPortfolioRepository) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	return m.skill, m.skillErr
}

func (m *MockPortfolioRepository) CreateSkill(ctx context.Context, skill *model.Skill) error {
	if m.createErr != nil {
		return m.createErr
	}
	skill.ID = 1
	return nil
}

func (m *MockPortfolioRepository) UpdateSkill(ctx context.Context, skill *model.Skill) error {
	return m.updateErr
}

func (m *MockPortfolioRepository) DeleteSkill(ctx context.Context, id int64) error {
	return m.deleteErr
}

// Project operations
func (m *MockPortfolioRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return m.projects, m.projectErr
}

func (m *MockPortfolioRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	return m.project, m.projectErr
}

func (m *MockPortfolioRepository) CreateProject(ctx context.Context, project *model.Project) error {
	if m.createErr != nil {
		return m.createErr
	}
	project.ID = 1
	return nil
}

func (m *MockPortfolioRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	return m.updateErr
}

func (m *MockPortfolioRepository) DeleteProject(ctx context.Context, id int64) error {
	return m.deleteErr
}

// Publication operations
func (m *MockPortfolioRepository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	return m.publications, m.publicationErr
}

func (m *MockPortfolioRepository) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	return m.publication, m.publicationErr
}

func (m *MockPortfolioRepository) CreatePublication(ctx context.Context, pub *model.Publication) error {
	if m.createErr != nil {
		return m.createErr
	}
	pub.ID = 1
	return nil
}

func (m *MockPortfolioRepository) UpdatePublication(ctx context.Context, pub *model.Publication) error {
	return m.updateErr
}

func (m *MockPortfolioRepository) DeletePublication(ctx context.Context, id int64) error {
	return m.deleteErr
}

// Portfolio operations
func (m *MockPortfolioRepository) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	return m.portfolioData, m.portfolioErr
}

// Test Profile Operations
func TestPortfolioService_GetProfile(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		profile: &model.Profile{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	profile, err := service.GetProfile(ctx)
	if err != nil {
		t.Errorf("GetProfile() error = %v, want nil", err)
	}
	if profile == nil {
		t.Error("GetProfile() returned nil profile")
	}
	if profile != nil && profile.Name != "John Doe" {
		t.Errorf("GetProfile() Name = %v, want John Doe", profile.Name)
	}
}

func TestPortfolioService_GetProfile_Error(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		profileErr: errors.New("database error"),
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetProfile(ctx)
	if err == nil {
		t.Error("GetProfile() expected error, got nil")
	}
}

func TestPortfolioService_CreateProfile(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	profile, err := service.CreateProfile(ctx, req)
	if err != nil {
		t.Errorf("CreateProfile() error = %v, want nil", err)
	}
	if profile == nil {
		t.Error("CreateProfile() returned nil profile")
	}
}

func TestPortfolioService_CreateProfile_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "",
		Email: "john@example.com",
	}

	_, err := service.CreateProfile(ctx, req)
	if err == nil {
		t.Error("CreateProfile() expected validation error, got nil")
	}
}

func TestPortfolioService_CreateProfile_RepositoryError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		createErr: errors.New("create failed"),
	}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	_, err := service.CreateProfile(ctx, req)
	if err == nil {
		t.Error("CreateProfile() expected error, got nil")
	}
}

func TestPortfolioService_UpdateProfile(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}

	profile, err := service.UpdateProfile(ctx, 1, req)
	if err != nil {
		t.Errorf("UpdateProfile() error = %v, want nil", err)
	}
	if profile == nil {
		t.Error("UpdateProfile() returned nil profile")
	}
	if profile != nil && profile.Name != "John Updated" {
		t.Errorf("UpdateProfile() Name = %v, want John Updated", profile.Name)
	}
}

func TestPortfolioService_UpdateProfile_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "",
		Email: "",
	}

	_, err := service.UpdateProfile(ctx, 1, req)
	if err == nil {
		t.Error("UpdateProfile() expected validation error, got nil")
	}
}

// Test Experience Operations
func TestPortfolioService_GetAllExperiences(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		experiences: []model.Experience{
			{ID: 1, Title: "Developer", Organization: "Tech Corp"},
			{ID: 2, Title: "Intern", Organization: "Startup Inc"},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	experiences, err := service.GetAllExperiences(ctx)
	if err != nil {
		t.Errorf("GetAllExperiences() error = %v, want nil", err)
	}
	if len(experiences) != 2 {
		t.Errorf("GetAllExperiences() returned %d experiences, want 2", len(experiences))
	}
}

func TestPortfolioService_GetExperienceByID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		experience: &model.Experience{ID: 1, Title: "Developer"},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	exp, err := service.GetExperienceByID(ctx, 1)
	if err != nil {
		t.Errorf("GetExperienceByID() error = %v, want nil", err)
	}
	if exp == nil {
		t.Error("GetExperienceByID() returned nil")
	}
}

func TestPortfolioService_GetExperienceByID_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetExperienceByID(ctx, 0)
	if err == nil {
		t.Error("GetExperienceByID() expected error for invalid ID, got nil")
	}

	_, err = service.GetExperienceByID(ctx, -1)
	if err == nil {
		t.Error("GetExperienceByID() expected error for negative ID, got nil")
	}
}

func TestPortfolioService_CreateExperience(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Developer",
		Organization: "Tech Corp",
		Type:         "work",
	}

	exp, err := service.CreateExperience(ctx, req)
	if err != nil {
		t.Errorf("CreateExperience() error = %v, want nil", err)
	}
	if exp == nil {
		t.Error("CreateExperience() returned nil")
	}
	if exp != nil && exp.Color != "cyan" {
		t.Errorf("CreateExperience() Color = %v, want cyan for work type", exp.Color)
	}
}

func TestPortfolioService_CreateExperience_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "",
		Organization: "Tech Corp",
		Type:         "work",
	}

	_, err := service.CreateExperience(ctx, req)
	if err == nil {
		t.Error("CreateExperience() expected validation error, got nil")
	}
}

func TestPortfolioService_UpdateExperience(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Senior Developer",
		Organization: "Tech Corp",
		Type:         "work",
	}

	exp, err := service.UpdateExperience(ctx, 1, req)
	if err != nil {
		t.Errorf("UpdateExperience() error = %v, want nil", err)
	}
	if exp == nil {
		t.Error("UpdateExperience() returned nil")
	}
}

func TestPortfolioService_UpdateExperience_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Developer",
		Organization: "Tech Corp",
		Type:         "work",
	}

	_, err := service.UpdateExperience(ctx, 0, req)
	if err == nil {
		t.Error("UpdateExperience() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_DeleteExperience(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteExperience(ctx, 1)
	if err != nil {
		t.Errorf("DeleteExperience() error = %v, want nil", err)
	}
}

func TestPortfolioService_DeleteExperience_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteExperience(ctx, 0)
	if err == nil {
		t.Error("DeleteExperience() expected error for invalid ID, got nil")
	}
}

// Test Skill Operations
func TestPortfolioService_GetAllSkills(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		skills: []model.Skill{
			{ID: 1, Name: "Go", Category: "Programming"},
			{ID: 2, Name: "Docker", Category: "DevOps"},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	skills, err := service.GetAllSkills(ctx)
	if err != nil {
		t.Errorf("GetAllSkills() error = %v, want nil", err)
	}
	if len(skills) != 2 {
		t.Errorf("GetAllSkills() returned %d skills, want 2", len(skills))
	}
}

func TestPortfolioService_GetSkillsByCategory(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		skills: []model.Skill{
			{ID: 1, Name: "Go", Category: "Programming"},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	skills, err := service.GetSkillsByCategory(ctx, "Programming")
	if err != nil {
		t.Errorf("GetSkillsByCategory() error = %v, want nil", err)
	}
	if len(skills) != 1 {
		t.Errorf("GetSkillsByCategory() returned %d skills, want 1", len(skills))
	}
}

func TestPortfolioService_GetSkillsByCategory_EmptyCategory(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetSkillsByCategory(ctx, "")
	if err == nil {
		t.Error("GetSkillsByCategory() expected error for empty category, got nil")
	}
}

func TestPortfolioService_GetSkillByID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		skill: &model.Skill{ID: 1, Name: "Go"},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	skill, err := service.GetSkillByID(ctx, 1)
	if err != nil {
		t.Errorf("GetSkillByID() error = %v, want nil", err)
	}
	if skill == nil {
		t.Error("GetSkillByID() returned nil")
	}
}

func TestPortfolioService_GetSkillByID_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetSkillByID(ctx, 0)
	if err == nil {
		t.Error("GetSkillByID() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_CreateSkill(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.SkillRequest{
		Category: "Programming",
		Name:     "Go",
		Level:    "advanced",
	}

	skill, err := service.CreateSkill(ctx, req)
	if err != nil {
		t.Errorf("CreateSkill() error = %v, want nil", err)
	}
	if skill == nil {
		t.Error("CreateSkill() returned nil")
	}
	if skill != nil && skill.Color != "black" {
		t.Errorf("CreateSkill() Color = %v, want black for advanced level", skill.Color)
	}
}

func TestPortfolioService_CreateSkill_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.SkillRequest{
		Category: "",
		Name:     "Go",
		Level:    "advanced",
	}

	_, err := service.CreateSkill(ctx, req)
	if err == nil {
		t.Error("CreateSkill() expected validation error, got nil")
	}
}

func TestPortfolioService_UpdateSkill(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.SkillRequest{
		Category: "Programming",
		Name:     "Golang",
		Level:    "advanced",
	}

	skill, err := service.UpdateSkill(ctx, 1, req)
	if err != nil {
		t.Errorf("UpdateSkill() error = %v, want nil", err)
	}
	if skill == nil {
		t.Error("UpdateSkill() returned nil")
	}
}

func TestPortfolioService_UpdateSkill_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.SkillRequest{
		Category: "Programming",
		Name:     "Go",
		Level:    "advanced",
	}

	_, err := service.UpdateSkill(ctx, 0, req)
	if err == nil {
		t.Error("UpdateSkill() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_DeleteSkill(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteSkill(ctx, 1)
	if err != nil {
		t.Errorf("DeleteSkill() error = %v, want nil", err)
	}
}

func TestPortfolioService_DeleteSkill_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteSkill(ctx, 0)
	if err == nil {
		t.Error("DeleteSkill() expected error for invalid ID, got nil")
	}
}

// Test Project Operations
func TestPortfolioService_GetAllProjects(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		projects: []model.Project{
			{ID: 1, Title: "Project 1"},
			{ID: 2, Title: "Project 2"},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	projects, err := service.GetAllProjects(ctx)
	if err != nil {
		t.Errorf("GetAllProjects() error = %v, want nil", err)
	}
	if len(projects) != 2 {
		t.Errorf("GetAllProjects() returned %d projects, want 2", len(projects))
	}
}

func TestPortfolioService_GetProjectByID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		project: &model.Project{ID: 1, Title: "Project 1"},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	project, err := service.GetProjectByID(ctx, 1)
	if err != nil {
		t.Errorf("GetProjectByID() error = %v, want nil", err)
	}
	if project == nil {
		t.Error("GetProjectByID() returned nil")
	}
}

func TestPortfolioService_GetProjectByID_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetProjectByID(ctx, 0)
	if err == nil {
		t.Error("GetProjectByID() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_CreateProject(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "New Project",
		Description: "A cool project",
	}

	project, err := service.CreateProject(ctx, req)
	if err != nil {
		t.Errorf("CreateProject() error = %v, want nil", err)
	}
	if project == nil {
		t.Error("CreateProject() returned nil")
	}
	if project != nil && project.Color != "cyan" {
		t.Errorf("CreateProject() Color = %v, want cyan as default", project.Color)
	}
}

func TestPortfolioService_CreateProject_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "",
		Description: "A cool project",
	}

	_, err := service.CreateProject(ctx, req)
	if err == nil {
		t.Error("CreateProject() expected validation error, got nil")
	}
}

func TestPortfolioService_UpdateProject(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "Updated Project",
		Description: "An updated project",
	}

	project, err := service.UpdateProject(ctx, 1, req)
	if err != nil {
		t.Errorf("UpdateProject() error = %v, want nil", err)
	}
	if project == nil {
		t.Error("UpdateProject() returned nil")
	}
}

func TestPortfolioService_UpdateProject_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "Project",
		Description: "A project",
	}

	_, err := service.UpdateProject(ctx, 0, req)
	if err == nil {
		t.Error("UpdateProject() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_DeleteProject(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteProject(ctx, 1)
	if err != nil {
		t.Errorf("DeleteProject() error = %v, want nil", err)
	}
}

func TestPortfolioService_DeleteProject_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeleteProject(ctx, 0)
	if err == nil {
		t.Error("DeleteProject() expected error for invalid ID, got nil")
	}
}

// Test Publication Operations
func TestPortfolioService_GetAllPublications(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		publications: []model.Publication{
			{ID: 1, Title: "Publication 1"},
			{ID: 2, Title: "Publication 2"},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	publications, err := service.GetAllPublications(ctx)
	if err != nil {
		t.Errorf("GetAllPublications() error = %v, want nil", err)
	}
	if len(publications) != 2 {
		t.Errorf("GetAllPublications() returned %d publications, want 2", len(publications))
	}
}

func TestPortfolioService_GetPublicationByID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		publication: &model.Publication{ID: 1, Title: "Publication 1"},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	pub, err := service.GetPublicationByID(ctx, 1)
	if err != nil {
		t.Errorf("GetPublicationByID() error = %v, want nil", err)
	}
	if pub == nil {
		t.Error("GetPublicationByID() returned nil")
	}
}

func TestPortfolioService_GetPublicationByID_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetPublicationByID(ctx, 0)
	if err == nil {
		t.Error("GetPublicationByID() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_CreatePublication(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title: "New Publication",
		Year:  2024,
	}

	pub, err := service.CreatePublication(ctx, req)
	if err != nil {
		t.Errorf("CreatePublication() error = %v, want nil", err)
	}
	if pub == nil {
		t.Error("CreatePublication() returned nil")
	}
	if pub != nil && pub.Color != "red" {
		t.Errorf("CreatePublication() Color = %v, want red as default", pub.Color)
	}
}

func TestPortfolioService_CreatePublication_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title: "",
		Year:  2024,
	}

	_, err := service.CreatePublication(ctx, req)
	if err == nil {
		t.Error("CreatePublication() expected validation error, got nil")
	}
}

func TestPortfolioService_UpdatePublication(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title: "Updated Publication",
		Year:  2024,
	}

	pub, err := service.UpdatePublication(ctx, 1, req)
	if err != nil {
		t.Errorf("UpdatePublication() error = %v, want nil", err)
	}
	if pub == nil {
		t.Error("UpdatePublication() returned nil")
	}
}

func TestPortfolioService_UpdatePublication_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title: "Publication",
		Year:  2024,
	}

	_, err := service.UpdatePublication(ctx, 0, req)
	if err == nil {
		t.Error("UpdatePublication() expected error for invalid ID, got nil")
	}
}

func TestPortfolioService_DeletePublication(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeletePublication(ctx, 1)
	if err != nil {
		t.Errorf("DeletePublication() error = %v, want nil", err)
	}
}

func TestPortfolioService_DeletePublication_InvalidID(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	err := service.DeletePublication(ctx, 0)
	if err == nil {
		t.Error("DeletePublication() expected error for invalid ID, got nil")
	}
}

// Test GetPortfolioData
func TestPortfolioService_GetPortfolioData(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		portfolioData: &model.PortfolioData{
			Profile: model.Profile{ID: 1, Name: "John Doe"},
			Experiences: []model.Experience{
				{ID: 1, Title: "Developer"},
			},
			Skills: map[string][]model.Skill{
				"Programming": {{ID: 1, Name: "Go"}},
			},
			Projects: []model.Project{
				{ID: 1, Title: "Project 1"},
			},
			Publications: []model.Publication{
				{ID: 1, Title: "Publication 1"},
			},
		},
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	data, err := service.GetPortfolioData(ctx)
	if err != nil {
		t.Errorf("GetPortfolioData() error = %v, want nil", err)
	}
	if data == nil {
		t.Error("GetPortfolioData() returned nil")
	}
	if data != nil && data.Profile.ID == 0 {
		t.Error("GetPortfolioData() Profile is empty")
	}
}

func TestPortfolioService_GetPortfolioData_Error(t *testing.T) {
	mockRepo := &MockPortfolioRepository{
		portfolioErr: errors.New("database error"),
	}

	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	_, err := service.GetPortfolioData(ctx)
	if err == nil {
		t.Error("GetPortfolioData() expected error, got nil")
	}
}

// Test SubmitContact
func TestPortfolioService_SubmitContact(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ContactRequest{
		Name:    "John Doe",
		Email:   "john@example.com",
		Subject: "Hello",
		Message: "This is a message",
	}

	err := service.SubmitContact(ctx, req)
	if err != nil {
		t.Errorf("SubmitContact() error = %v, want nil", err)
	}
}

func TestPortfolioService_SubmitContact_ValidationError(t *testing.T) {
	mockRepo := &MockPortfolioRepository{}
	service := NewPortfolioService(mockRepo)
	ctx := context.Background()

	req := &dto.ContactRequest{
		Name:    "",
		Email:   "john@example.com",
		Subject: "Hello",
		Message: "This is a message",
	}

	err := service.SubmitContact(ctx, req)
	if err == nil {
		t.Error("SubmitContact() expected validation error, got nil")
	}
}

// Test Helper Functions
func TestGetColorForType(t *testing.T) {
	tests := []struct {
		name         string
		expType      string
		defaultColor string
		want         string
	}{
		{"work type", "work", "", "cyan"},
		{"internship type", "internship", "", "pink"},
		{"campus type", "campus", "", "yellow"},
		{"competition type", "competition", "", "purple"},
		{"unknown type", "unknown", "", "gray"},
		{"with default color", "work", "red", "red"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getColorForType(tt.expType, tt.defaultColor)
			if result != tt.want {
				t.Errorf("getColorForType() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestGetColorForLevel(t *testing.T) {
	tests := []struct {
		name         string
		level        string
		defaultColor string
		want         string
	}{
		{"advanced level", "advanced", "", "black"},
		{"intermediate level", "intermediate", "", "gray"},
		{"beginner level", "beginner", "", "white"},
		{"unknown level", "unknown", "", "gray"},
		{"with default color", "advanced", "blue", "blue"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getColorForLevel(tt.level, tt.defaultColor)
			if result != tt.want {
				t.Errorf("getColorForLevel() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestGetDefaultColor(t *testing.T) {
	tests := []struct {
		name         string
		color        string
		defaultColor string
		want         string
	}{
		{"with color", "blue", "red", "blue"},
		{"empty color", "", "red", "red"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDefaultColor(tt.color, tt.defaultColor)
			if result != tt.want {
				t.Errorf("getDefaultColor() = %v, want %v", result, tt.want)
			}
		})
	}
}
