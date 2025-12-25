package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// SkillRepositoryInterface defines the interface for skill repository
type SkillRepositoryInterface interface {
	GetAllSkills(ctx context.Context) ([]model.Skill, error)
	GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*model.Skill, error)
	CreateSkill(ctx context.Context, skill *model.Skill) error
	UpdateSkill(ctx context.Context, skill *model.Skill) error
	DeleteSkill(ctx context.Context, id int64) error
}

// SkillRepository implements SkillRepositoryInterface
type SkillRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewSkillRepository creates a new skill repository
func NewSkillRepository(db database.PgxIface, log *zap.Logger) SkillRepositoryInterface {
	return &SkillRepository{
		db:  db,
		log: log,
	}
}

// GetAllSkills retrieves all skills
func (r *SkillRepository) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
	query := `SELECT id, category, name, COALESCE(level, 'intermediate'), COALESCE(color, 'gray') 
		FROM skills ORDER BY category, name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.log.Error("Failed to get skills", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var skills []model.Skill
	for rows.Next() {
		var skill model.Skill
		err := rows.Scan(&skill.ID, &skill.Category, &skill.Name, &skill.Level, &skill.Color)
		if err != nil {
			r.log.Error("Failed to scan skill", zap.Error(err))
			continue
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

// GetSkillsByCategory retrieves skills by category
func (r *SkillRepository) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
	query := `SELECT id, category, name, COALESCE(level, 'intermediate'), COALESCE(color, 'gray') 
		FROM skills WHERE category = $1 ORDER BY name`

	rows, err := r.db.Query(ctx, query, category)
	if err != nil {
		r.log.Error("Failed to get skills by category", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var skills []model.Skill
	for rows.Next() {
		var skill model.Skill
		err := rows.Scan(&skill.ID, &skill.Category, &skill.Name, &skill.Level, &skill.Color)
		if err != nil {
			r.log.Error("Failed to scan skill", zap.Error(err))
			continue
		}
		skills = append(skills, skill)
	}
	return skills, nil
}

// GetSkillByID retrieves a skill by ID
func (r *SkillRepository) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
	query := `SELECT id, category, name, COALESCE(level, 'intermediate'), COALESCE(color, 'gray') 
		FROM skills WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)
	var skill model.Skill
	err := row.Scan(&skill.ID, &skill.Category, &skill.Name, &skill.Level, &skill.Color)
	if err != nil {
		r.log.Error("Failed to get skill by ID", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return &skill, nil
}

// CreateSkill creates a new skill
func (r *SkillRepository) CreateSkill(ctx context.Context, skill *model.Skill) error {
	query := `INSERT INTO skills (category, name, level, color) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(ctx, query, skill.Category, skill.Name, skill.Level, skill.Color)
	err := row.Scan(&skill.ID)
	if err != nil {
		r.log.Error("Failed to create skill", zap.Error(err))
		return err
	}
	return nil
}

// UpdateSkill updates a skill
func (r *SkillRepository) UpdateSkill(ctx context.Context, skill *model.Skill) error {
	query := `UPDATE skills SET category = $1, name = $2, level = $3, color = $4 WHERE id = $5`

	_, err := r.db.Exec(ctx, query, skill.Category, skill.Name, skill.Level, skill.Color, skill.ID)
	if err != nil {
		r.log.Error("Failed to update skill", zap.Error(err))
		return err
	}
	return nil
}

// DeleteSkill deletes a skill
func (r *SkillRepository) DeleteSkill(ctx context.Context, id int64) error {
	query := `DELETE FROM skills WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete skill", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}
