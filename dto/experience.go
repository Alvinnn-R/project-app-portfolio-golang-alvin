package dto

// ExperienceRequest represents the request body for creating/updating experience
type ExperienceRequest struct {
	Title        string `json:"title"`
	Organization string `json:"organization"`
	Period       string `json:"period"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Color        string `json:"color"`
}
