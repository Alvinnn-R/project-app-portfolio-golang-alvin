package service

import (
	"errors"
	"regexp"
	"session-19/dto"
	"strings"
)

// Validation errors
var (
	ErrNameRequired         = errors.New("name is required")
	ErrEmailRequired        = errors.New("email is required")
	ErrEmailInvalid         = errors.New("email format is invalid")
	ErrTitleRequired        = errors.New("title is required")
	ErrOrganizationRequired = errors.New("organization is required")
	ErrCategoryRequired     = errors.New("category is required")
	ErrSkillNameRequired    = errors.New("skill name is required")
	ErrLevelRequired        = errors.New("level is required")
	ErrDescriptionRequired  = errors.New("description is required")
	ErrMessageRequired      = errors.New("message is required")
	ErrAuthorsRequired      = errors.New("authors is required")
	ErrJournalRequired      = errors.New("journal is required")
	ErrYearRequired         = errors.New("year is required")
	ErrYearInvalid          = errors.New("year must be between 1900 and 2100")
	ErrInvalidID            = errors.New("invalid ID")
)

// emailRegex is a simple regex for email validation
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateProfileRequest validates a profile request
func ValidateProfileRequest(req *dto.ProfileRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return ErrNameRequired
	}
	if strings.TrimSpace(req.Email) == "" {
		return ErrEmailRequired
	}
	if !emailRegex.MatchString(req.Email) {
		return ErrEmailInvalid
	}
	return nil
}

// ValidateExperienceRequest validates an experience request
func ValidateExperienceRequest(req *dto.ExperienceRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return ErrTitleRequired
	}
	if strings.TrimSpace(req.Organization) == "" {
		return ErrOrganizationRequired
	}
	return nil
}

// ValidateSkillRequest validates a skill request
func ValidateSkillRequest(req *dto.SkillRequest) error {
	if strings.TrimSpace(req.Category) == "" {
		return ErrCategoryRequired
	}
	if strings.TrimSpace(req.Name) == "" {
		return ErrSkillNameRequired
	}
	if strings.TrimSpace(req.Level) == "" {
		return ErrLevelRequired
	}
	return nil
}

// ValidateProjectRequest validates a project request
func ValidateProjectRequest(req *dto.ProjectRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return ErrTitleRequired
	}
	if strings.TrimSpace(req.Description) == "" {
		return ErrDescriptionRequired
	}
	return nil
}

// ValidatePublicationRequest validates a publication request
func ValidatePublicationRequest(req *dto.PublicationRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return ErrTitleRequired
	}
	if strings.TrimSpace(req.Authors) == "" {
		return ErrAuthorsRequired
	}
	if strings.TrimSpace(req.Journal) == "" {
		return ErrJournalRequired
	}
	if req.Year == 0 {
		return ErrYearRequired
	}
	if req.Year < 1900 || req.Year > 2100 {
		return ErrYearInvalid
	}
	return nil
}

// ValidateContactRequest validates a contact form request
func ValidateContactRequest(req *dto.ContactRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return ErrNameRequired
	}
	if strings.TrimSpace(req.Email) == "" {
		return ErrEmailRequired
	}
	if !emailRegex.MatchString(req.Email) {
		return ErrEmailInvalid
	}
	if strings.TrimSpace(req.Message) == "" {
		return ErrMessageRequired
	}
	return nil
}

// ValidateID validates that an ID is valid (positive)
func ValidateID(id int64) error {
	if id <= 0 {
		return ErrInvalidID
	}
	return nil
}

// ValidateCategory validates that a category is not empty
func ValidateCategory(category string) error {
	if strings.TrimSpace(category) == "" {
		return ErrCategoryRequired
	}
	return nil
}
