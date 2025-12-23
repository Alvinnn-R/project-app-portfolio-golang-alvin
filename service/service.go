package service

import "session-19/repository"

// Service contains all services
type Service struct {
	PortfolioService PortfolioServiceInterface
	AuthService      AuthServiceInterface
}

// NewService creates a new service with all sub-services
func NewService(repo repository.Repository) Service {
	return Service{
		PortfolioService: NewPortfolioService(repo.PortfolioRepo),
		AuthService:      NewAuthService(repo.UserRepo),
	}
}
