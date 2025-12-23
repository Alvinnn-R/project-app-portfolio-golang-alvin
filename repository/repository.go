package repository

import (
	"session-19/database"

	"go.uber.org/zap"
)

// Repository contains all repositories
type Repository struct {
	PortfolioRepo PortfolioRepositoryInterface
	UserRepo      UserRepositoryInterface
}

// NewRepository creates a new repository with all sub-repositories
func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		PortfolioRepo: NewPortfolioRepository(db, log),
		UserRepo:      NewUserRepository(db, log),
	}
}
