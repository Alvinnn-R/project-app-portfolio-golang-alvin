package model

import "time"

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
