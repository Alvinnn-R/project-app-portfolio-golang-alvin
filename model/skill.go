package model

// Skill represents a skill with category and level
type Skill struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Level    string `json:"level"` // beginner, intermediate, advanced
	Color    string `json:"color"`
}
