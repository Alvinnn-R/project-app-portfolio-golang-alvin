package repository

import (
	"session-19/database"

	"go.uber.org/zap"
)

// Repository contains all repositories
type Repository struct {
	PortfolioRepo PortfolioRepositoryInterface
}

// NewRepository creates a new repository with all sub-repositories
func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		PortfolioRepo: NewPortfolioRepository(db, log),
	}
}
