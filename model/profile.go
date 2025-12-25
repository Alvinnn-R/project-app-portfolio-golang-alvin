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
