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

// newTestPublicationRepository creates a new test publication repository
func newTestPublicationRepository() (*PublicationRepository, *database.MockDB) {
	mockDB := new(database.MockDB)
	logger := zap.NewNop()
	repo := NewPublicationRepository(mockDB, logger)
	return repo.(*PublicationRepository), mockDB
}

// ==================== Publication Repository Tests ====================

func TestPublicationRepository_GetAllPublications_Success(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	now := time.Now()
	mockRows := database.NewMockRows([][]any{
		{int64(1), "Publication 1", "Author 1, Author 2", "Journal 1", 2024, "Description 1", "/image1.jpg", "https://pub1.com", "red", now},
		{int64(2), "Publication 2", "Author 3", "Journal 2", 2023, "Description 2", "/image2.jpg", "https://pub2.com", "blue", now},
	})
	mockRows.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		data := mockRows.Data[mockRows.CurrentIndex]
		*dest[0].(*int64) = data[0].(int64)
		*dest[1].(*string) = data[1].(string)
		*dest[2].(*string) = data[2].(string)
		*dest[3].(*string) = data[3].(string)
		*dest[4].(*int) = data[4].(int)
		*dest[5].(*string) = data[5].(string)
		*dest[6].(*string) = data[6].(string)
		*dest[7].(*string) = data[7].(string)
		*dest[8].(*string) = data[8].(string)
		*dest[9].(*time.Time) = data[9].(time.Time)
	}).Return(nil)
	mockRows.On("Close").Return()
	mockRows.On("Err").Return(nil)

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRows, nil).Once()

	publications, err := repo.GetAllPublications(ctx)

	assert.NoError(t, err)
	assert.Len(t, publications, 2)
	assert.Equal(t, "Publication 1", publications[0].Title)
	assert.Equal(t, "Publication 2", publications[1].Title)
	mockDB.AssertExpectations(t)
}

func TestPublicationRepository_GetAllPublications_QueryError(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(nil, errors.New("query failed")).Once()

	publications, err := repo.GetAllPublications(ctx)

	assert.Error(t, err)
	assert.Nil(t, publications)
	mockDB.AssertExpectations(t)
}

func TestPublicationRepository_GetPublicationByID_Success(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	now := time.Now()
	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*string) = "Publication 1"
		*dest[2].(*string) = "Author 1, Author 2"
		*dest[3].(*string) = "Journal 1"
		*dest[4].(*int) = 2024
		*dest[5].(*string) = "Description 1"
		*dest[6].(*string) = "/image1.jpg"
		*dest[7].(*string) = "https://pub1.com"
		*dest[8].(*string) = "red"
		*dest[9].(*time.Time) = now
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	publication, err := repo.GetPublicationByID(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, publication)
	assert.Equal(t, "Publication 1", publication.Title)
	assert.Equal(t, 2024, publication.Year)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestPublicationRepository_GetPublicationByID_NotFound(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("no rows")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	publication, err := repo.GetPublicationByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, publication)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestPublicationRepository_CreatePublication_Success(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	now := time.Now()
	publication := &model.Publication{
		Title:          "New Publication",
		Authors:        "Author 1, Author 2",
		Journal:        "Journal Name",
		Year:           2024,
		Description:    "Publication description",
		ImageURL:       "/image.jpg",
		PublicationURL: "https://publication.com",
		Color:          "red",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*time.Time) = now
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreatePublication(ctx, publication)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), publication.ID)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestPublicationRepository_CreatePublication_Error(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	publication := &model.Publication{
		Title:   "New Publication",
		Authors: "Author 1",
		Journal: "Journal Name",
		Year:    2024,
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("insert failed")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreatePublication(ctx, publication)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestPublicationRepository_UpdatePublication_Success(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	publication := &model.Publication{
		ID:          1,
		Title:       "Updated Publication",
		Authors:     "Author 1, Author 2",
		Journal:     "Updated Journal",
		Year:        2024,
		Description: "Updated description",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.UpdatePublication(ctx, publication)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestPublicationRepository_UpdatePublication_Error(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	publication := &model.Publication{
		ID:      1,
		Title:   "Updated Publication",
		Authors: "Author 1",
		Journal: "Updated Journal",
		Year:    2024,
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("update failed")).Once()

	err := repo.UpdatePublication(ctx, publication)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestPublicationRepository_DeletePublication_Success(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.DeletePublication(ctx, 1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestPublicationRepository_DeletePublication_Error(t *testing.T) {
	repo, mockDB := newTestPublicationRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("delete failed")).Once()

	err := repo.DeletePublication(ctx, 1)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
