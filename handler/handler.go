package handler

import (
	"html/template"
	"session-19/service"

	"go.uber.org/zap"
)

// Handler contains all handlers for the application
type Handler struct {
	PortfolioHandler *PortfolioHandler
	AuthHandler      *AuthHandler
	AdminHandler     *AdminHandler
}

// NewHandler creates a new handler with all sub-handlers
func NewHandler(svc service.Service, log *zap.Logger, tmpl *template.Template) Handler {
	return Handler{
		PortfolioHandler: NewPortfolioHandler(svc.PortfolioService, log),
		AuthHandler:      NewAuthHandler(svc.AuthService, log, tmpl),
		AdminHandler:     NewAdminHandler(svc.PortfolioService, log, tmpl),
	}
}
