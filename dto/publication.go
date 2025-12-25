package dto

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
