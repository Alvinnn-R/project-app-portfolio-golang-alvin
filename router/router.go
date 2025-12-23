package router

import (
	"net/http"
	"session-19/handler"
	mCostume "session-19/middleware"
	"session-19/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// NewRouter creates a new router with all routes configured
func NewRouter(h handler.Handler, svc service.Service, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Custom middleware
	mw := mCostume.NewMiddlewareCustome(svc, log)
	r.Use(mw.Logging)

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	// Main portfolio page (HTML template)
	r.Get("/", h.PortfolioHandler.RenderPortfolio)

	// Auth routes (public)
	r.Get("/login", h.AuthHandler.LoginView)
	r.Post("/login", h.AuthHandler.Login)
	r.Get("/logout", h.AuthHandler.LogoutView)
	r.Post("/logout", h.AuthHandler.Logout)
	r.Get("/page401", h.AuthHandler.Page401)

	// Admin routes (protected)
	r.Route("/admin", func(r chi.Router) {
		r.Use(mCostume.AuthMiddleware)

		// Dashboard
		r.Get("/", http.RedirectHandler("/admin/dashboard", http.StatusSeeOther).ServeHTTP)
		r.Get("/dashboard", h.AdminHandler.Dashboard)

		// Profile
		r.Get("/profile", h.AdminHandler.ProfileEdit)
		r.Post("/profile/save", h.AdminHandler.ProfileSave)

		// Experiences
		r.Get("/experiences", h.AdminHandler.ExperiencesList)
		r.Get("/experiences/new", h.AdminHandler.ExperienceForm)
		r.Get("/experiences/edit/{id}", h.AdminHandler.ExperienceForm)
		r.Post("/experiences/save", h.AdminHandler.ExperienceSave)
		r.Post("/experiences/delete/{id}", h.AdminHandler.ExperienceDelete)

		// Skills
		r.Get("/skills", h.AdminHandler.SkillsList)
		r.Get("/skills/new", h.AdminHandler.SkillForm)
		r.Get("/skills/edit/{id}", h.AdminHandler.SkillForm)
		r.Post("/skills/save", h.AdminHandler.SkillSave)
		r.Post("/skills/delete/{id}", h.AdminHandler.SkillDelete)

		// Projects
		r.Get("/projects", h.AdminHandler.ProjectsList)
		r.Get("/projects/new", h.AdminHandler.ProjectForm)
		r.Get("/projects/edit/{id}", h.AdminHandler.ProjectForm)
		r.Post("/projects/save", h.AdminHandler.ProjectSave)
		r.Post("/projects/delete/{id}", h.AdminHandler.ProjectDelete)

		// Publications
		r.Get("/publications", h.AdminHandler.PublicationsList)
		r.Get("/publications/new", h.AdminHandler.PublicationForm)
		r.Get("/publications/edit/{id}", h.AdminHandler.PublicationForm)
		r.Post("/publications/save", h.AdminHandler.PublicationSave)
		r.Post("/publications/delete/{id}", h.AdminHandler.PublicationDelete)
	})

	// API v1 routes
	r.Mount("/api/v1", ApiV1Routes(h, mw))

	return r
}

// ApiV1Routes creates API v1 routes
func ApiV1Routes(h handler.Handler, mw mCostume.MiddlewareCostume) *chi.Mux {
	r := chi.NewRouter()

	// Portfolio data endpoint (JSON)
	r.Get("/portfolio", h.PortfolioHandler.GetPortfolioData)

	// Profile routes
	r.Route("/profile", func(r chi.Router) {
		r.Get("/", h.PortfolioHandler.GetProfile)
		r.Post("/", h.PortfolioHandler.CreateProfile)
		r.Put("/{id}", h.PortfolioHandler.UpdateProfile)
	})

	// Experience routes
	r.Route("/experiences", func(r chi.Router) {
		r.Get("/", h.PortfolioHandler.GetAllExperiences)
		r.Post("/", h.PortfolioHandler.CreateExperience)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.PortfolioHandler.GetExperienceByID)
			r.Put("/", h.PortfolioHandler.UpdateExperience)
			r.Delete("/", h.PortfolioHandler.DeleteExperience)
		})
	})

	// Skill routes
	r.Route("/skills", func(r chi.Router) {
		r.Get("/", h.PortfolioHandler.GetAllSkills)
		r.Post("/", h.PortfolioHandler.CreateSkill)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.PortfolioHandler.GetSkillByID)
			r.Put("/", h.PortfolioHandler.UpdateSkill)
			r.Delete("/", h.PortfolioHandler.DeleteSkill)
		})
	})

	// Project routes
	r.Route("/projects", func(r chi.Router) {
		r.Get("/", h.PortfolioHandler.GetAllProjects)
		r.Post("/", h.PortfolioHandler.CreateProject)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.PortfolioHandler.GetProjectByID)
			r.Put("/", h.PortfolioHandler.UpdateProject)
			r.Delete("/", h.PortfolioHandler.DeleteProject)
		})
	})

	// Publication routes
	r.Route("/publications", func(r chi.Router) {
		r.Get("/", h.PortfolioHandler.GetAllPublications)
		r.Post("/", h.PortfolioHandler.CreatePublication)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.PortfolioHandler.GetPublicationByID)
			r.Put("/", h.PortfolioHandler.UpdatePublication)
			r.Delete("/", h.PortfolioHandler.DeletePublication)
		})
	})

	// Contact form submission
	r.Post("/contact", h.PortfolioHandler.SubmitContact)

	return r
}
