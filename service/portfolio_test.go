package service

import (
	"context"
	"errors"
	"session-19/dto"
	"session-19/model"
	"session-19/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// newTestService creates a new test portfolio service with mock repository
func newTestService() (PortfolioServiceInterface, *repository.MockPortfolioRepository) {
	mockRepo := new(repository.MockPortfolioRepository)
	service := NewPortfolioService(mockRepo)
	return service, mockRepo
}

// ==================== Profile Service Tests ====================

func TestPortfolioService_GetProfile_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.Profile{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	mockRepo.On("GetProfile", ctx).Return(expected, nil).Once()

	result, err := svc.GetProfile(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetProfile_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetProfile", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetProfile(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateProfile_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Title: "Software Engineer",
	}
	mockRepo.On("CreateProfile", ctx, mock.AnythingOfType("*model.Profile")).Return(nil).Once()

	result, err := svc.CreateProfile(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John Doe", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateProfile_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Title: "Software Engineer",
	}
	mockRepo.On("CreateProfile", ctx, mock.AnythingOfType("*model.Profile")).Return(errors.New("insert failed")).Once()

	result, err := svc.CreateProfile(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateProfile_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "Updated Name",
		Email: "updated@example.com",
		Title: "Senior Engineer",
	}
	mockRepo.On("UpdateProfile", ctx, mock.AnythingOfType("*model.Profile")).Return(nil).Once()

	result, err := svc.UpdateProfile(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Updated Name", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateProfile_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProfileRequest{
		Name:  "Updated Name",
		Email: "updated@example.com",
		Title: "Senior Engineer",
	}
	mockRepo.On("UpdateProfile", ctx, mock.AnythingOfType("*model.Profile")).Return(errors.New("update failed")).Once()

	result, err := svc.UpdateProfile(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ==================== Experience Service Tests ====================

func TestPortfolioService_GetAllExperiences_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := []model.Experience{
		{ID: 1, Title: "Software Engineer"},
		{ID: 2, Title: "Tech Lead"},
	}
	mockRepo.On("GetAllExperiences", ctx).Return(expected, nil).Once()

	result, err := svc.GetAllExperiences(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetAllExperiences_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetAllExperiences", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetAllExperiences(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetExperienceByID_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.Experience{ID: 1, Title: "Software Engineer"}
	mockRepo.On("GetExperienceByID", ctx, int64(1)).Return(expected, nil).Once()

	result, err := svc.GetExperienceByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetExperienceByID_NotFound(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetExperienceByID", ctx, int64(999)).Return(nil, errors.New("not found")).Once()

	result, err := svc.GetExperienceByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateExperience_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Software Engineer",
		Organization: "Tech Corp",
		Description:  "Building software",
	}
	mockRepo.On("CreateExperience", ctx, mock.AnythingOfType("*model.Experience")).Return(nil).Once()

	result, err := svc.CreateExperience(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateExperience_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Software Engineer",
		Organization: "Tech Corp",
		Description:  "Building software",
	}
	mockRepo.On("CreateExperience", ctx, mock.AnythingOfType("*model.Experience")).Return(errors.New("insert failed")).Once()

	result, err := svc.CreateExperience(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateExperience_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Senior Engineer",
		Organization: "Tech Corp",
		Description:  "Leading teams",
	}
	mockRepo.On("UpdateExperience", ctx, mock.AnythingOfType("*model.Experience")).Return(nil).Once()

	result, err := svc.UpdateExperience(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateExperience_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ExperienceRequest{
		Title:        "Senior Engineer",
		Organization: "Tech Corp",
		Description:  "Leading teams",
	}
	mockRepo.On("UpdateExperience", ctx, mock.AnythingOfType("*model.Experience")).Return(errors.New("update failed")).Once()

	result, err := svc.UpdateExperience(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_DeleteExperience_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("DeleteExperience", ctx, int64(1)).Return(nil).Once()

	err := svc.DeleteExperience(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_DeleteExperience_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("DeleteExperience", ctx, int64(1)).Return(errors.New("delete failed")).Once()

	err := svc.DeleteExperience(ctx, 1)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// ==================== Skill Service Tests ====================

func TestPortfolioService_GetAllSkills_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := []model.Skill{
		{ID: 1, Name: "Go"},
		{ID: 2, Name: "Python"},
	}
	mockRepo.On("GetAllSkills", ctx).Return(expected, nil).Once()

	result, err := svc.GetAllSkills(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetAllSkills_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetAllSkills", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetAllSkills(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetSkillsByCategory_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := []model.Skill{
		{ID: 1, Name: "Go", Category: "Backend"},
	}
	mockRepo.On("GetSkillsByCategory", ctx, "Backend").Return(expected, nil).Once()

	result, err := svc.GetSkillsByCategory(ctx, "Backend")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetSkillByID_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.Skill{ID: 1, Name: "Go"}
	mockRepo.On("GetSkillByID", ctx, int64(1)).Return(expected, nil).Once()

	result, err := svc.GetSkillByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateSkill_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.SkillRequest{
		Name:     "Go",
		Category: "Backend",
	}
	mockRepo.On("CreateSkill", ctx, mock.AnythingOfType("*model.Skill")).Return(nil).Once()

	result, err := svc.CreateSkill(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateSkill_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.SkillRequest{
		Name:     "Golang",
		Category: "Backend",
	}
	mockRepo.On("UpdateSkill", ctx, mock.AnythingOfType("*model.Skill")).Return(nil).Once()

	result, err := svc.UpdateSkill(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_DeleteSkill_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("DeleteSkill", ctx, int64(1)).Return(nil).Once()

	err := svc.DeleteSkill(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// ==================== Project Service Tests ====================

func TestPortfolioService_GetAllProjects_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := []model.Project{
		{ID: 1, Title: "Project 1"},
		{ID: 2, Title: "Project 2"},
	}
	mockRepo.On("GetAllProjects", ctx).Return(expected, nil).Once()

	result, err := svc.GetAllProjects(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetAllProjects_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetAllProjects", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetAllProjects(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetProjectByID_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.Project{ID: 1, Title: "Project 1"}
	mockRepo.On("GetProjectByID", ctx, int64(1)).Return(expected, nil).Once()

	result, err := svc.GetProjectByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreateProject_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "New Project",
		Description: "Project description",
	}
	mockRepo.On("CreateProject", ctx, mock.AnythingOfType("*model.Project")).Return(nil).Once()

	result, err := svc.CreateProject(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdateProject_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.ProjectRequest{
		Title:       "Updated Project",
		Description: "Updated description",
	}
	mockRepo.On("UpdateProject", ctx, mock.AnythingOfType("*model.Project")).Return(nil).Once()

	result, err := svc.UpdateProject(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_DeleteProject_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("DeleteProject", ctx, int64(1)).Return(nil).Once()

	err := svc.DeleteProject(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// ==================== Publication Service Tests ====================

func TestPortfolioService_GetAllPublications_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := []model.Publication{
		{ID: 1, Title: "Publication 1", Authors: "Author 1", Journal: "Journal 1"},
		{ID: 2, Title: "Publication 2", Authors: "Author 2", Journal: "Journal 2"},
	}
	mockRepo.On("GetAllPublications", ctx).Return(expected, nil).Once()

	result, err := svc.GetAllPublications(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetAllPublications_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetAllPublications", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetAllPublications(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetPublicationByID_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.Publication{ID: 1, Title: "Publication 1", Authors: "Author 1", Journal: "Journal 1"}
	mockRepo.On("GetPublicationByID", ctx, int64(1)).Return(expected, nil).Once()

	result, err := svc.GetPublicationByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_CreatePublication_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title:   "New Publication",
		Authors: "Author Name",
		Journal: "Journal Name",
		Year:    2024,
	}
	mockRepo.On("CreatePublication", ctx, mock.AnythingOfType("*model.Publication")).Return(nil).Once()

	result, err := svc.CreatePublication(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_UpdatePublication_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	req := &dto.PublicationRequest{
		Title:   "Updated Publication",
		Authors: "Updated Author",
		Journal: "Updated Journal",
		Year:    2024,
	}
	mockRepo.On("UpdatePublication", ctx, mock.AnythingOfType("*model.Publication")).Return(nil).Once()

	result, err := svc.UpdatePublication(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_DeletePublication_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("DeletePublication", ctx, int64(1)).Return(nil).Once()

	err := svc.DeletePublication(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// ==================== Portfolio Data Tests ====================

func TestPortfolioService_GetPortfolioData_Success(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	expected := &model.PortfolioData{
		Profile: model.Profile{Name: "John Doe"},
	}
	mockRepo.On("GetPortfolioData", ctx).Return(expected, nil).Once()

	result, err := svc.GetPortfolioData(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestPortfolioService_GetPortfolioData_Error(t *testing.T) {
	svc, mockRepo := newTestService()
	ctx := context.Background()

	mockRepo.On("GetPortfolioData", ctx).Return(nil, errors.New("database error")).Once()

	result, err := svc.GetPortfolioData(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ==================== Contact Service Tests ====================

func TestPortfolioService_SubmitContact_Success(t *testing.T) {
	svc, _ := newTestService()
	ctx := context.Background()

	req := &dto.ContactRequest{
		Name:    "Visitor",
		Email:   "visitor@example.com",
		Message: "Hello!",
	}

	err := svc.SubmitContact(ctx, req)

	assert.NoError(t, err)
}
