package repository

import (
	"context"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// PortfolioRepositoryInterface defines the interface for portfolio repository
type PortfolioRepositoryInterface interface {
	// Profile operations
	GetProfile(ctx context.Context) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error

	// Experience operations
	GetAllExperiences(ctx context.Context) ([]model.Experience, error)
	GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error)
	CreateExperience(ctx context.Context, exp *model.Experience) error
	UpdateExperience(ctx context.Context, exp *model.Experience) error
	DeleteExperience(ctx context.Context, id int64) error

	// Skill operations
	GetAllSkills(ctx context.Context) ([]model.Skill, error)
	GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error)
	GetSkillByID(ctx context.Context, id int64) (*model.Skill, error)
	CreateSkill(ctx context.Context, skill *model.Skill) error
	UpdateSkill(ctx context.Context, skill *model.Skill) error
	DeleteSkill(ctx context.Context, id int64) error

	// Project operations
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) error
	UpdateProject(ctx context.Context, project *model.Project) error
	DeleteProject(ctx context.Context, id int64) error

	// Publication operations
	GetAllPublications(ctx context.Context) ([]model.Publication, error)
	GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error)
	CreatePublication(ctx context.Context, pub *model.Publication) error
	UpdatePublication(ctx context.Context, pub *model.Publication) error
	DeletePublication(ctx context.Context, id int64) error

	// Full portfolio data
	GetPortfolioData(ctx context.Context) (*model.PortfolioData, error)
}

// PortfolioRepository implements PortfolioRepositoryInterface
type PortfolioRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewPortfolioRepository creates a new portfolio repository
func NewPortfolioRepository(db database.PgxIface, log *zap.Logger) PortfolioRepositoryInterface {
	return &PortfolioRepository{
		db:  db,
		log: log,
	}
}

// GetProfile retrieves the main profile
func (r *PortfolioRepository) GetProfile(ctx context.Context) (*model.Profile, error) {
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
func (r *PortfolioRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
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
func (r *PortfolioRepository) UpdateProfile(ctx context.Context, profile *model.Profile) error {
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

// GetAllExperiences retrieves all experiences
func (r *PortfolioRepository) GetAllExperiences(ctx context.Context) ([]model.Experience, error) {
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
func (r *PortfolioRepository) GetExperienceByID(ctx context.Context, id int64) (*model.Experience, error) {
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
func (r *PortfolioRepository) CreateExperience(ctx context.Context, exp *model.Experience) error {
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
func (r *PortfolioRepository) UpdateExperience(ctx context.Context, exp *model.Experience) error {
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
func (r *PortfolioRepository) DeleteExperience(ctx context.Context, id int64) error {
	query := `DELETE FROM experiences WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete experience", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}

// GetAllSkills retrieves all skills
func (r *PortfolioRepository) GetAllSkills(ctx context.Context) ([]model.Skill, error) {
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
func (r *PortfolioRepository) GetSkillsByCategory(ctx context.Context, category string) ([]model.Skill, error) {
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
func (r *PortfolioRepository) GetSkillByID(ctx context.Context, id int64) (*model.Skill, error) {
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
func (r *PortfolioRepository) CreateSkill(ctx context.Context, skill *model.Skill) error {
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
func (r *PortfolioRepository) UpdateSkill(ctx context.Context, skill *model.Skill) error {
	query := `UPDATE skills SET category = $1, name = $2, level = $3, color = $4 WHERE id = $5`

	_, err := r.db.Exec(ctx, query, skill.Category, skill.Name, skill.Level, skill.Color, skill.ID)
	if err != nil {
		r.log.Error("Failed to update skill", zap.Error(err))
		return err
	}
	return nil
}

// DeleteSkill deletes a skill
func (r *PortfolioRepository) DeleteSkill(ctx context.Context, id int64) error {
	query := `DELETE FROM skills WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete skill", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}

// GetAllProjects retrieves all projects
func (r *PortfolioRepository) GetAllProjects(ctx context.Context) ([]model.Project, error) {
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
func (r *PortfolioRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
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
func (r *PortfolioRepository) CreateProject(ctx context.Context, project *model.Project) error {
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
func (r *PortfolioRepository) UpdateProject(ctx context.Context, project *model.Project) error {
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
func (r *PortfolioRepository) DeleteProject(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete project", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}

// GetAllPublications retrieves all publications
func (r *PortfolioRepository) GetAllPublications(ctx context.Context) ([]model.Publication, error) {
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
func (r *PortfolioRepository) GetPublicationByID(ctx context.Context, id int64) (*model.Publication, error) {
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
func (r *PortfolioRepository) CreatePublication(ctx context.Context, pub *model.Publication) error {
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
func (r *PortfolioRepository) UpdatePublication(ctx context.Context, pub *model.Publication) error {
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
func (r *PortfolioRepository) DeletePublication(ctx context.Context, id int64) error {
	query := `DELETE FROM publications WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		r.log.Error("Failed to delete publication", zap.Error(err), zap.Int64("id", id))
		return err
	}
	return nil
}

// GetPortfolioData retrieves all portfolio data in one call
func (r *PortfolioRepository) GetPortfolioData(ctx context.Context) (*model.PortfolioData, error) {
	data := &model.PortfolioData{}

	// Get profile
	profile, err := r.GetProfile(ctx)
	if err != nil {
		r.log.Warn("No profile found, using empty profile", zap.Error(err))
		data.Profile = model.Profile{}
	} else {
		data.Profile = *profile
	}

	// Get experiences
	experiences, err := r.GetAllExperiences(ctx)
	if err != nil {
		r.log.Warn("Failed to get experiences", zap.Error(err))
		data.Experiences = []model.Experience{}
	} else {
		data.Experiences = experiences
	}

	// Get skills and group by category
	skills, err := r.GetAllSkills(ctx)
	if err != nil {
		r.log.Warn("Failed to get skills", zap.Error(err))
		data.Skills = make(map[string][]model.Skill)
	} else {
		skillMap := make(map[string][]model.Skill)
		for _, skill := range skills {
			skillMap[skill.Category] = append(skillMap[skill.Category], skill)
		}
		data.Skills = skillMap
	}

	// Get projects
	projects, err := r.GetAllProjects(ctx)
	if err != nil {
		r.log.Warn("Failed to get projects", zap.Error(err))
		data.Projects = []model.Project{}
	} else {
		data.Projects = projects
	}

	// Get publications
	publications, err := r.GetAllPublications(ctx)
	if err != nil {
		r.log.Warn("Failed to get publications", zap.Error(err))
		data.Publications = []model.Publication{}
	} else {
		data.Publications = publications
	}

	return data, nil
}
