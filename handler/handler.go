package handler

import (
	"session-19/service"

	"go.uber.org/zap"
)

// Handler contains all handlers for the application
type Handler struct {
	PortfolioHandler *PortfolioHandler
}

// NewHandler creates a new handler with all sub-handlers
func NewHandler(svc service.Service, log *zap.Logger) Handler {
	return Handler{
		PortfolioHandler: NewPortfolioHandler(svc.PortfolioService, log),
	}
}
