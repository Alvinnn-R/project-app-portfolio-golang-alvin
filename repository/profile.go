package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// ProfileRepositoryInterface defines the interface for profile repository
type ProfileRepositoryInterface interface {
	GetProfile(ctx context.Context) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error
}

// ProfileRepository implements ProfileRepositoryInterface
type ProfileRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewProfileRepository creates a new profile repository
func NewProfileRepository(db database.PgxIface, log *zap.Logger) ProfileRepositoryInterface {
	return &ProfileRepository{
		db:  db,
		log: log,
	}
}

// GetProfile retrieves the main profile
func (r *ProfileRepository) GetProfile(ctx context.Context) (*model.Profile, error) {
	query := `SELECT id, name, COALESCE(title, ''), COALESCE(description, ''), 
		COALESCE(photo_url, ''), email, COALESCE(linkedin_url, ''), 
		COALESCE(github_url, ''), COALESCE(cv_url, ''), created_at, updated_at 
		FROM profile LIMIT 1`

	row := r.db.QueryRow(ctx, query)
	var p model.Profile
	err := row.Scan(&p.ID, &p.Name, &p.Title, &p.Description, &p.PhotoURL,
		&p.Email, &p.LinkedInURL, &p.GithubURL, &p.CVURL, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		r.log.Error("Failed to get profile", zap.Error(err))
		return nil, err
	}
	return &p, nil
}

// CreateProfile creates a new profile
func (r *ProfileRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	query := `INSERT INTO profile (name, title, description, photo_url, email, linkedin_url, github_url, cv_url) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at, updated_at`

	row := r.db.QueryRow(ctx, query, profile.Name, profile.Title, profile.Description,
		profile.PhotoURL, profile.Email, profile.LinkedInURL, profile.GithubURL, profile.CVURL)

	err := row.Scan(&profile.ID, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		r.log.Error("Failed to create profile", zap.Error(err))
		return err
	}
	return nil
}

// UpdateProfile updates the profile
func (r *ProfileRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	query := `UPDATE profile SET name = $1, title = $2, description = $3, photo_url = $4, 
		email = $5, linkedin_url = $6, github_url = $7, cv_url = $8, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $9`

	_, err := r.db.Exec(ctx, query, profile.Name, profile.Title, profile.Description,
		profile.PhotoURL, profile.Email, profile.LinkedInURL, profile.GithubURL, profile.CVURL, profile.ID)
	if err != nil {
		r.log.Error("Failed to update profile", zap.Error(err))
		return err
	}
	return nil
}
