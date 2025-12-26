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

// newTestExperienceRepository creates a new test experience repository
func newTestExperienceRepository() (*ExperienceRepository, *database.MockDB) {
	mockDB := new(database.MockDB)
	logger := zap.NewNop()
	repo := NewExperienceRepository(mockDB, logger)
	return repo.(*ExperienceRepository), mockDB
}

// ==================== Experience Repository Tests ====================

func TestExperienceRepository_GetAllExperiences_Success(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockRows := database.NewMockRows([][]any{
		{int64(1), "Developer", "Tech Corp", "2023-Present", "Description", "work", "cyan", time.Now()},
		{int64(2), "Intern", "Startup", "2022-2023", "Internship", "internship", "pink", time.Now()},
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
		*dest[7].(*time.Time) = data[7].(time.Time)
	}).Return(nil)
	mockRows.On("Close").Return()
	mockRows.On("Err").Return(nil)

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRows, nil).Once()

	experiences, err := repo.GetAllExperiences(ctx)

	assert.NoError(t, err)
	assert.Len(t, experiences, 2)
	assert.Equal(t, "Developer", experiences[0].Title)
	assert.Equal(t, "Intern", experiences[1].Title)
	mockDB.AssertExpectations(t)
}

func TestExperienceRepository_GetAllExperiences_QueryError(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(nil, errors.New("query failed")).Once()

	experiences, err := repo.GetAllExperiences(ctx)

	assert.Error(t, err)
	assert.Nil(t, experiences)
	mockDB.AssertExpectations(t)
}

func TestExperienceRepository_GetExperienceByID_Success(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*string) = "Developer"
		*dest[2].(*string) = "Tech Corp"
		*dest[3].(*string) = "2023-Present"
		*dest[4].(*string) = "Description"
		*dest[5].(*string) = "work"
		*dest[6].(*string) = "cyan"
		*dest[7].(*time.Time) = time.Now()
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	exp, err := repo.GetExperienceByID(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, exp)
	assert.Equal(t, int64(1), exp.ID)
	assert.Equal(t, "Developer", exp.Title)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestExperienceRepository_GetExperienceByID_NotFound(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("no rows")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	exp, err := repo.GetExperienceByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, exp)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestExperienceRepository_CreateExperience_Success(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	exp := &model.Experience{
		Title:        "Developer",
		Organization: "Tech Corp",
		Type:         "work",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*time.Time) = time.Now()
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateExperience(ctx, exp)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), exp.ID)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestExperienceRepository_CreateExperience_Error(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	exp := &model.Experience{
		Title:        "Developer",
		Organization: "Tech Corp",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("insert failed")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateExperience(ctx, exp)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestExperienceRepository_UpdateExperience_Success(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	exp := &model.Experience{
		ID:           1,
		Title:        "Senior Developer",
		Organization: "Tech Corp",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.UpdateExperience(ctx, exp)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestExperienceRepository_UpdateExperience_Error(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	exp := &model.Experience{
		ID:           1,
		Title:        "Senior Developer",
		Organization: "Tech Corp",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("update failed")).Once()

	err := repo.UpdateExperience(ctx, exp)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestExperienceRepository_DeleteExperience_Success(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.DeleteExperience(ctx, 1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestExperienceRepository_DeleteExperience_Error(t *testing.T) {
	repo, mockDB := newTestExperienceRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("delete failed")).Once()

	err := repo.DeleteExperience(ctx, 1)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
