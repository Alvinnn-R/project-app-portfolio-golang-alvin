package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// ProjectRepositoryInterface defines the interface for project repository
type ProjectRepositoryInterface interface {
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) error
	UpdateProject(ctx context.Context, project *model.Project) error
	DeleteProject(ctx context.Context, id int64) error
}

// ProjectRepository implements ProjectRepositoryInterface
type ProjectRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewProjectRepository creates a new project repository
func NewProjectRepository(db database.PgxIface, log *zap.Logger) ProjectRepositoryInterface {
	return &ProjectRepository{
		db:  db,
		log: log,
	}
}

// GetAllProjects retrieves all projects
func (r *ProjectRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	query := `SELECT id, title, COALESCE(description, ''), COALESCE(image_url, ''), 
		COALESCE(project_url, ''), COALESCE(github_url, ''), COALESCE(tech_stack, ''), 
		COALESCE(color, 'cyan'), COALESCE(profile_id, 0), created_at 
		FROM projects ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.log.Error("Failed to get projects", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.ProjectURL,
			&p.GithubURL, &p.TechStack, &p.Color, &p.ProfileID, &p.CreatedAt)
		if err != nil {
			r.log.Error("Failed to scan project", zap.Error(err))
			continue
		}
		projects = append(projects, p)
	}
	return projects, nil
}

// GetProjectByID retrieves a project by ID
func (r *ProjectRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	query := `SELECT id, title, COALESCE(description, ''), COALESCE(image_url, ''), 
		COALESCE(project_url, ''), COALESCE(github_url, ''), COALESCE(tech_stack, ''), 
		COALESCE(color, 'cyan'), COALESCE(profile_id, 0), created_at 
		FROM projects WHERE id = $1`

	row := r.db.QueryRow(ctx, query, id)
	var p model.Project
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.ProjectURL,
		&p.GithubURL, &p.TechStack, &p.Color, &p.ProfileID, &p.CreatedAt)
	if err != nil {
		r.log.Error("Failed to get project by ID", zap.Error(err), zap.Int64("id", id))
		return nil, err
	}
	return &p, nil
}

// CreateProject creates a new project
func (r *ProjectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	query := `INSERT INTO projects (title, description, image_url, project_url, github_url, tech_stack, color, profile_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, created_at`

	row := r.db.QueryRow(ctx, query, project.Title, project.Description, project.ImageURL,
		project.ProjectURL, project.GithubURL, project.TechStack, project.Color, project.ProfileID)

	err := row.Scan(&project.ID, &project.CreatedAt)
	if err != nil {
		r.log.Error("Failed to create project", zap.Error(err))
		return err
	}
	return nil
}

// UpdateProject updates a project
func (r *ProjectRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	query := `UPDATE projects SET title = $1, description = $2, image_url = $3, project_url = $4, 
		github_url = $5, tech_stack = $6, color = $7, profile_id = $8 WHERE id = $9`

	_, err := r.db.Exec(ctx, query, project.Title, project.Description, project.ImageURL,
		project.ProjectURL, project.GithubURL, project.TechStack, project.Color, project.ProfileID, project.ID)
	if err != nil {
		r.log.Error("Failed to update project", zap.Error(err))
		return err
	}
	return nil
}

// DeleteProject deletes a project
func (r *ProjectRepository) DeleteProject(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete project", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}
