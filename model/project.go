package model

import "time"

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
