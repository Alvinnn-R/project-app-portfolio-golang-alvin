package repository

import (
	"context"
	"errors"
	"session-19/database"
	"session-19/model"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// newTestProjectRepository creates a new test project repository
func newTestProjectRepository() (*ProjectRepository, *database.MockDB) {
	mockDB := new(database.MockDB)
	logger := zap.NewNop()
	repo := NewProjectRepository(mockDB, logger)
	return repo.(*ProjectRepository), mockDB
}

// ==================== Project Repository Tests ====================

func TestProjectRepository_GetAllProjects_Success(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	now := time.Now()
	mockRows := database.NewMockRows([][]any{
		{int64(1), "Project 1", "Description 1", "/image1.jpg", "https://project1.com", "https://github.com/project1", "Go, React", "cyan", int64(1), now},
		{int64(2), "Project 2", "Description 2", "/image2.jpg", "https://project2.com", "https://github.com/project2", "Python, Vue", "blue", int64(1), now},
	})
	mockRows.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		data := mockRows.Data[mockRows.CurrentIndex]
		*dest[0].(*int64) = data[0].(int64)
		*dest[1].(*string) = data[1].(string)
		*dest[2].(*string) = data[2].(string)
		*dest[3].(*string) = data[3].(string)
		*dest[4].(*string) = data[4].(string)
		*dest[5].(*string) = data[5].(string)
		*dest[6].(*string) = data[6].(string)
		*dest[7].(*string) = data[7].(string)
		*dest[8].(*int64) = data[8].(int64)
		*dest[9].(*time.Time) = data[9].(time.Time)
	}).Return(nil)
	mockRows.On("Close").Return()
	mockRows.On("Err").Return(nil)

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRows, nil).Once()

	projects, err := repo.GetAllProjects(ctx)

	assert.NoError(t, err)
	assert.Len(t, projects, 2)
	assert.Equal(t, "Project 1", projects[0].Title)
	assert.Equal(t, "Project 2", projects[1].Title)
	mockDB.AssertExpectations(t)
}

func TestProjectRepository_GetAllProjects_QueryError(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(nil, errors.New("query failed")).Once()

	projects, err := repo.GetAllProjects(ctx)

	assert.Error(t, err)
	assert.Nil(t, projects)
	mockDB.AssertExpectations(t)
}

func TestProjectRepository_GetProjectByID_Success(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	now := time.Now()
	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*string) = "Project 1"
		*dest[2].(*string) = "Description 1"
		*dest[3].(*string) = "/image1.jpg"
		*dest[4].(*string) = "https://project1.com"
		*dest[5].(*string) = "https://github.com/project1"
		*dest[6].(*string) = "Go, React"
		*dest[7].(*string) = "cyan"
		*dest[8].(*int64) = 1
		*dest[9].(*time.Time) = now
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	project, err := repo.GetProjectByID(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, "Project 1", project.Title)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProjectRepository_GetProjectByID_NotFound(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("no rows")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	project, err := repo.GetProjectByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, project)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProjectRepository_CreateProject_Success(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	now := time.Now()
	project := &model.Project{
		Title:       "New Project",
		Description: "New Description",
		ImageURL:    "/image.jpg",
		ProjectURL:  "https://project.com",
		GithubURL:   "https://github.com/project",
		TechStack:   "Go, React",
		Color:       "cyan",
		ProfileID:   1,
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*time.Time) = now
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateProject(ctx, project)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), project.ID)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProjectRepository_CreateProject_Error(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	project := &model.Project{
		Title:       "New Project",
		Description: "New Description",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("insert failed")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateProject(ctx, project)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProjectRepository_UpdateProject_Success(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	project := &model.Project{
		ID:          1,
		Title:       "Updated Project",
		Description: "Updated Description",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.UpdateProject(ctx, project)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestProjectRepository_UpdateProject_Error(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	project := &model.Project{
		ID:          1,
		Title:       "Updated Project",
		Description: "Updated Description",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("update failed")).Once()

	err := repo.UpdateProject(ctx, project)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestProjectRepository_DeleteProject_Success(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.DeleteProject(ctx, 1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestProjectRepository_DeleteProject_Error(t *testing.T) {
	repo, mockDB := newTestProjectRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("delete failed")).Once()

	err := repo.DeleteProject(ctx, 1)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
