package dto

// SkillRequest represents the request body for creating/updating skill
type SkillRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Color    string `json:"color"`
}
