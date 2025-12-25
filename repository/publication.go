package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// PublicationRepositoryInterface defines the interface for publication repository
type PublicationRepositoryInterface interface {
	GetAllPublications(ctx context.Context) ([]model.Publication, error)
	GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error)
	CreatePublication(ctx context.Context, pub *model.Publication) error
	UpdatePublication(ctx context.Context, pub *model.Publication) error
	DeletePublication(ctx context.Context, id int64) error
}

// PublicationRepository implements PublicationRepositoryInterface
type PublicationRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewPublicationRepository creates a new publication repository
func NewPublicationRepository(db database.PgxIface, log *zap.Logger) PublicationRepositoryInterface {
	return &PublicationRepository{
		db:  db,
		log: log,
	}
}

// GetAllPublications retrieves all publications
func (r *PublicationRepository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
	query := `SELECT id, title, COALESCE(authors, ''), COALESCE(journal, ''), COALESCE(year, 0), 
		COALESCE(description, ''), COALESCE(image_url, ''), COALESCE(publication_url, ''), 
		COALESCE(color, 'red'), created_at FROM publications ORDER BY year DESC, created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.log.Error("Failed to get publications", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var publications []model.Publication
	for rows.Next() {
		var p model.Publication
		err := rows.Scan(&p.ID, &p.Title, &p.Authors, &p.Journal, &p.Year,
			&p.Description, &p.ImageURL, &p.PublicationURL, &p.Color, &p.CreatedAt)
		if err != nil {
			r.log.Error("Failed to scan publication", zap.Error(err))
			continue
		}
		publications = append(publications, p)
	}
	return publications, nil
}

// GetPublicationByID retrieves a publication by ID
func (r *PublicationRepository) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
	query := `SELECT id, title, COALESCE(authors, ''), COALESCE(journal, ''), COALESCE(year, 0), 
		COALESCE(description, ''), COALESCE(image_url, ''), COALESCE(publication_url, ''), 
		COALESCE(color, 'red'), created_at FROM publications WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)
	var p model.Publication
	err := row.Scan(&p.ID, &p.Title, &p.Authors, &p.Journal, &p.Year,
		&p.Description, &p.ImageURL, &p.PublicationURL, &p.Color, &p.CreatedAt)
	if err != nil {
		r.log.Error("Failed to get publication by ID", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return &p, nil
}

// CreatePublication creates a new publication
func (r *PublicationRepository) CreatePublication(ctx context.Context, pub *model.Publication) error {
	query := `INSERT INTO publications (title, authors, journal, year, description, image_url, publication_url, color) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`

	row := r.db.QueryRow(ctx, query, pub.Title, pub.Authors, pub.Journal, pub.Year,
		pub.Description, pub.ImageURL, pub.PublicationURL, pub.Color)

	err := row.Scan(&pub.ID, &pub.CreatedAt)
	if err != nil {
		r.log.Error("Failed to create publication", zap.Error(err))
		return err
	}
	return nil
}

// UpdatePublication updates a publication
func (r *PublicationRepository) UpdatePublication(ctx context.Context, pub *model.Publication) error {
	query := `UPDATE publications SET title = $1, authors = $2, journal = $3, year = $4, 
		description = $5, image_url = $6, publication_url = $7, color = $8 WHERE id = $9`

	_, err := r.db.Exec(ctx, query, pub.Title, pub.Authors, pub.Journal, pub.Year,
		pub.Description, pub.ImageURL, pub.PublicationURL, pub.Color, pub.ID)
	if err != nil {
		r.log.Error("Failed to update publication", zap.Error(err))
		return err
	}
	return nil
}

// DeletePublication deletes a publication
func (r *PublicationRepository) DeletePublication(ctx context.Context, id int64) error {
	query := `DELETE FROM publications WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete publication", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}
