package model

import "time"

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
