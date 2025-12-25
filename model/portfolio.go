package model

// PortfolioData represents all data needed for the portfolio page
type PortfolioData struct {
	Profile      Profile            `json:"profile"`
	Experiences  []Experience       `json:"experiences"`
	Skills       map[string][]Skill `json:"skills"`
	Projects     []Project          `json:"projects"`
	Publications []Publication      `json:"publications"`
}
