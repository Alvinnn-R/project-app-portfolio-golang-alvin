package repository

import (
	"context"
	"errors"
	"session-19/database"
	"session-19/model"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// newTestSkillRepository creates a new test skill repository
func newTestSkillRepository() (*SkillRepository, *database.MockDB) {
	mockDB := new(database.MockDB)
	logger := zap.NewNop()
	repo := NewSkillRepository(mockDB, logger)
	return repo.(*SkillRepository), mockDB
}

// ==================== Skill Repository Tests ====================

func TestSkillRepository_GetAllSkills_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockRows := database.NewMockRows([][]any{
		{int64(1), "Backend", "Go", "advanced", "black"},
		{int64(2), "Frontend", "React", "intermediate", "gray"},
	})
	mockRows.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		data := mockRows.Data[mockRows.CurrentIndex]
		*dest[0].(*int64) = data[0].(int64)
		*dest[1].(*string) = data[1].(string)
		*dest[2].(*string) = data[2].(string)
		*dest[3].(*string) = data[3].(string)
		*dest[4].(*string) = data[4].(string)
	}).Return(nil)
	mockRows.On("Close").Return()
	mockRows.On("Err").Return(nil)

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRows, nil).Once()

	skills, err := repo.GetAllSkills(ctx)

	assert.NoError(t, err)
	assert.Len(t, skills, 2)
	assert.Equal(t, "Go", skills[0].Name)
	assert.Equal(t, "React", skills[1].Name)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_GetAllSkills_QueryError(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(nil, errors.New("query failed")).Once()

	skills, err := repo.GetAllSkills(ctx)

	assert.Error(t, err)
	assert.Nil(t, skills)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_GetSkillsByCategory_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockRows := database.NewMockRows([][]any{
		{int64(1), "Backend", "Go", "advanced", "black"},
	})
	mockRows.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		data := mockRows.Data[mockRows.CurrentIndex]
		*dest[0].(*int64) = data[0].(int64)
		*dest[1].(*string) = data[1].(string)
		*dest[2].(*string) = data[2].(string)
		*dest[3].(*string) = data[3].(string)
		*dest[4].(*string) = data[4].(string)
	}).Return(nil)
	mockRows.On("Close").Return()
	mockRows.On("Err").Return(nil)

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRows, nil).Once()

	skills, err := repo.GetSkillsByCategory(ctx, "Backend")

	assert.NoError(t, err)
	assert.Len(t, skills, 1)
	assert.Equal(t, "Go", skills[0].Name)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_GetSkillsByCategory_QueryError(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockDB.On("Query", ctx, mock.AnythingOfType("string"), mock.Anything).Return(nil, errors.New("query failed")).Once()

	skills, err := repo.GetSkillsByCategory(ctx, "Backend")

	assert.Error(t, err)
	assert.Nil(t, skills)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_GetSkillByID_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
		*dest[1].(*string) = "Backend"
		*dest[2].(*string) = "Go"
		*dest[3].(*string) = "advanced"
		*dest[4].(*string) = "black"
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	skill, err := repo.GetSkillByID(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, skill)
	assert.Equal(t, "Go", skill.Name)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestSkillRepository_GetSkillByID_NotFound(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("no rows")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	skill, err := repo.GetSkillByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, skill)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestSkillRepository_CreateSkill_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	skill := &model.Skill{
		Name:     "Go",
		Category: "Backend",
		Level:    "advanced",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).([]any)
		*dest[0].(*int64) = 1
	}).Return(nil).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateSkill(ctx, skill)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), skill.ID)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestSkillRepository_CreateSkill_Error(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	skill := &model.Skill{
		Name:     "Go",
		Category: "Backend",
	}

	mockRow := new(database.MockRow)
	mockRow.On("Scan", mock.Anything).Return(errors.New("insert failed")).Once()

	mockDB.On("QueryRow", ctx, mock.AnythingOfType("string"), mock.Anything).Return(mockRow).Once()

	err := repo.CreateSkill(ctx, skill)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
	mockRow.AssertExpectations(t)
}

func TestSkillRepository_UpdateSkill_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	skill := &model.Skill{
		ID:       1,
		Name:     "Golang",
		Category: "Backend",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.UpdateSkill(ctx, skill)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_UpdateSkill_Error(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	skill := &model.Skill{
		ID:       1,
		Name:     "Golang",
		Category: "Backend",
	}

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("update failed")).Once()

	err := repo.UpdateSkill(ctx, skill)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_DeleteSkill_Success(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, nil).Once()

	err := repo.DeleteSkill(ctx, 1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestSkillRepository_DeleteSkill_Error(t *testing.T) {
	repo, mockDB := newTestSkillRepository()
	ctx := context.Background()

	mockDB.On("Exec", ctx, mock.AnythingOfType("string"), mock.Anything).Return(pgconn.CommandTag{}, errors.New("delete failed")).Once()

	err := repo.DeleteSkill(ctx, 1)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}
