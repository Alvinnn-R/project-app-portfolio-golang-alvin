package service

import (
	"context"
	"session-19/dto"
)

// ContactServiceInterface defines the interface for contact service
type ContactServiceInterface interface {
	SubmitContact(ctx context.Context, req *dto.ContactRequest) error
}

// ContactService implements ContactServiceInterface
type ContactService struct{}

// NewContactService creates a new contact service
func NewContactService() ContactServiceInterface {
	return &ContactService{}
}

// SubmitContact handles contact form submission
func (s *ContactService) SubmitContact(ctx context.Context, req *dto.ContactRequest) error {
	// In a real application, you would send an email or store the contact message
	// For now, we just return nil (success)
	return nil
}
