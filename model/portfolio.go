package model

import "time"

// Profile represents the main portfolio profile
type Profile struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PhotoURL    string    `json:"photo_url"`
	Email       string    `json:"email"`
	LinkedInURL string    `json:"linkedin_url"`
	GithubURL   string    `json:"github_url"`
	CVURL       string    `json:"cv_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Experience represents work experience, internship, campus activities, or competitions
type Experience struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Organization string    `json:"organization"`
	Period       string    `json:"period"`
	Description  string    `json:"description"`
	Type         string    `json:"type"` // work, internship, campus, competition
	Color        string    `json:"color"`
	CreatedAt    time.Time `json:"created_at"`
}

// Skill represents a skill with category and level
type Skill struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Level    string `json:"level"` // beginner, intermediate, advanced
	Color    string `json:"color"`
}

// Project represents a portfolio project
type Project struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ProjectURL  string    `json:"project_url"`
	GithubURL   string    `json:"github_url"`
	TechStack   string    `json:"tech_stack"`
	Color       string    `json:"color"`
	ProfileID   int64     `json:"profile_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// Publication represents academic or professional publications
type Publication struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	Authors        string    `json:"authors"`
	Journal        string    `json:"journal"`
	Year           int       `json:"year"`
	Description    string    `json:"description"`
	ImageURL       string    `json:"image_url"`
	PublicationURL string    `json:"publication_url"`
	Color          string    `json:"color"`
	CreatedAt      time.Time `json:"created_at"`
}

// PortfolioData represents all data needed for the portfolio page
type PortfolioData struct {
	Profile      Profile            `json:"profile"`
	Experiences  []Experience       `json:"experiences"`
	Skills       map[string][]Skill `json:"skills"`
	Projects     []Project          `json:"projects"`
	Publications []Publication      `json:"publications"`
}
