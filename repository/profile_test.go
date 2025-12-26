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

// newTestProfileRepository creates a new test profile repository
func newTestProfileRepository() (*ProfileRepository, *database.MockDB) {
	mockDB := new(database.MockDB)
	logger := zap.NewNop()
	repo := NewProfileRepository(mockDB, logger)
	return repo.(*ProfileRepository), mockDB
}

// ==================== Profile Repository Tests ====================

func TestProfileRepository_GetProfile_Success(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*string) = "John Doe"
		*dest[2].(*string) = "Software Engineer"
		*dest[3].(*string) = "Description"
		*dest[4].(*string) = "/photo.jpg"
		*dest[5].(*string) = "john@example.com"
		*dest[6].(*string) = "https://linkedin.com/in/john"
		*dest[7].(*string) = "https://github.com/john"
		*dest[8].(*string) = "/cv.pdf"
		*dest[9].(*time.Time) = time.Now()
		*dest[10].(*time.Time) = time.Now()
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	profile, err := repo.GetProfile(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, int64(1), profile.ID)
	assert.Equal(t, "John Doe", profile.Name)
	assert.Equal(t, "john@example.com", profile.Email)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProfileRepository_GetProfile_Error(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("no rows")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	profile, err := repo.GetProfile(ctx)

	assert.Error(t, err)
	assert.Nil(t, profile)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProfileRepository_CreateProfile_Success(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	profile := &model.Profile{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*time.Time) = time.Now()
		*dest[2].(*time.Time) = time.Now()
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateProfile(ctx, profile)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), profile.ID)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProfileRepository_CreateProfile_Error(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	profile := &model.Profile{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("insert failed")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateProfile(ctx, profile)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestProfileRepository_UpdateProfile_Success(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	profile := &model.Profile{
		ID:    1,
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.UpdateProfile(ctx, profile)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestProfileRepository_UpdateProfile_Error(t *testing.T) {
	repo, mockDB := newTestProfileRepository()
	ctx := context.Background()

	profile := &model.Profile{
		ID:    1,
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("update failed")).Once()

	err := repo.UpdateProfile(ctx, profile)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
