package dto

import "errors"

// ProfileRequest represents the request body for creating/updating a profile
type ProfileRequest struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PhotoURL    string `json:"photo_url"`
	Email       string `json:"email"`
	LinkedInURL string `json:"linkedin_url"`
	GithubURL   string `json:"github_url"`
	CVURL       string `json:"cv_url"`
}

// Validate validates the profile request
func (p *ProfileRequest) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if len(p.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	if p.Email == "" {
		return errors.New("email is required")
	}
	if len(p.Email) > 100 {
		return errors.New("email must be less than 100 characters")
	}
	return nil
}

// ExperienceRequest represents the request body for creating/updating experience
type ExperienceRequest struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
	Period       string `json:"period"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Color        string `json:"color"`
}

// Validate validates the experience request
func (e *ExperienceRequest) Validate() error {
	if e.Title == "" {
		return errors.New("title is required")
	}
	if len(e.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	if e.Organization == "" {
		return errors.New("organization is required")
	}
	if len(e.Organization) > 200 {
		return errors.New("organization must be less than 200 characters")
	}
	validTypes := map[string]bool{"work": true, "internship": true, "campus": true, "competition": true}
	if !validTypes[e.Type] {
		return errors.New("type must be one of: work, internship, campus, competition")
	}
	return nil
}

// SkillRequest represents the request body for creating/updating skill
type SkillRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Color    string `json:"color"`
}

// Validate validates the skill request
func (s *SkillRequest) Validate() error {
	if s.Category == "" {
		return errors.New("category is required")
	}
	if len(s.Category) > 100 {
		return errors.New("category must be less than 100 characters")
	}
	if s.Name == "" {
		return errors.New("name is required")
	}
	if len(s.Name) > 100 {
		return errors.New("name must be less than 100 characters")
	}
	validLevels := map[string]bool{"beginner": true, "intermediate": true, "advanced": true}
	if s.Level != "" && !validLevels[s.Level] {
		return errors.New("level must be one of: beginner, intermediate, advanced")
	}
	return nil
}

// ProjectRequest represents the request body for creating/updating project
type ProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ProjectURL  string `json:"project_url"`
	GithubURL   string `json:"github_url"`
	TechStack   string `json:"tech_stack"`
	Color       string `json:"color"`
	ProfileID   int64  `json:"profile_id"`
}

// Validate validates the project request
func (p *ProjectRequest) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if len(p.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	return nil
}

// PublicationRequest represents the request body for creating/updating publication
type PublicationRequest struct {
	Title          string `json:"title"`
	Authors        string `json:"authors"`
	Journal        string `json:"journal"`
	Year           int    `json:"year"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
	PublicationURL string `json:"publication_url"`
	Color          string `json:"color"`
}

// Validate validates the publication request
func (p *PublicationRequest) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if len(p.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	if p.Year < 1900 || p.Year > 2100 {
		return errors.New("year must be between 1900 and 2100")
	}
	return nil
}

// ContactRequest represents contact form submission
type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// Validate validates the contact request
func (c *ContactRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Message == "" {
		return errors.New("message is required")
	}
	return nil
}
