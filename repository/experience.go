package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// ExperienceRepositoryInterface defines the interface for experience repository
type ExperienceRepositoryInterface interface {
	GetAllExperiences(ctx context.Context) ([]model.Experience, error)
	GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error)
	CreateExperience(ctx context.Context, exp *model.Experience) error
	UpdateExperience(ctx context.Context, exp *model.Experience) error
	DeleteExperience(ctx context.Context, id int64) error
}

// ExperienceRepository implements ExperienceRepositoryInterface
type ExperienceRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewExperienceRepository creates a new experience repository
func NewExperienceRepository(db database.PgxIface, log *zap.Logger) ExperienceRepositoryInterface {
	return &ExperienceRepository{
		db:  db,
		log: log,
	}
}

// GetAllExperiences retrieves all experiences
func (r *ExperienceRepository) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
	query := `SELECT id, title, organization, COALESCE(period, ''), COALESCE(description, ''), 
		type, COALESCE(color, 'cyan'), created_at FROM experiences ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.log.Error("Failed to get experiences", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var experiences []model.Experience
	for rows.Next() {
		var exp model.Experience
		err := rows.Scan(&exp.ID, &exp.Title, &exp.Organization, &exp.Period,
			&exp.Description, &exp.Type, &exp.Color, &exp.CreatedAt)
		if err != nil {
			r.log.Error("Failed to scan experience", zap.Error(err))
			continue
		}
		experiences = append(experiences, exp)
	}
	return experiences, nil
}

// GetExperienceByID retrieves an experience by ID
func (r *ExperienceRepository) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
	query := `SELECT id, title, organization, COALESCE(period, ''), COALESCE(description, ''), 
		type, COALESCE(color, 'cyan'), created_at FROM experiences WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)
	var exp model.Experience
	err := row.Scan(&exp.ID, &exp.Title, &exp.Organization, &exp.Period,
		&exp.Description, &exp.Type, &exp.Color, &exp.CreatedAt)
	if err != nil {
		r.log.Error("Failed to get experience by ID", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return &exp, nil
}

// CreateExperience creates a new experience
func (r *ExperienceRepository) CreateExperience(ctx context.Context, exp *model.Experience) error {
	query := `INSERT INTO experiences (title, organization, period, description, type, color) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`

	row := r.db.QueryRow(ctx, query, exp.Title, exp.Organization, exp.Period,
		exp.Description, exp.Type, exp.Color)

	err := row.Scan(&exp.ID, &exp.CreatedAt)
	if err != nil {
		r.log.Error("Failed to create experience", zap.Error(err))
		return err
	}
	return nil
}

// UpdateExperience updates an experience
func (r *ExperienceRepository) UpdateExperience(ctx context.Context, exp *model.Experience) error {
	query := `UPDATE experiences SET title = $1, organization = $2, period = $3, 
		description = $4, type = $5, color = $6 WHERE id = $7`

	_, err := r.db.Exec(ctx, query, exp.Title, exp.Organization, exp.Period,
		exp.Description, exp.Type, exp.Color, exp.ID)
	if err != nil {
		r.log.Error("Failed to update experience", zap.Error(err))
		return err
	}
	return nil
}

// DeleteExperience deletes an experience
func (r *ExperienceRepository) DeleteExperience(ctx context.Context, id int64) error {
	query := `DELETE FROM experiences WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete experience", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}
