package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"session-19/dto"
	"session-19/service"
	"session-19/utils"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// PortfolioHandler handles HTTP requests for portfolio
type PortfolioHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
	tmpl    *template.Template
}

// NewPortfolioHandler creates a new portfolio handler
func NewPortfolioHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *PortfolioHandler {
	// Parse templates with custom functions
	funcMap := template.FuncMap{
		"split": func(s, sep string) []string {
			return strings.Split(s, sep)
		},
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"title":    strings.Title,
		"contains": strings.Contains,
		"formatYear": func(year int) string {
			if year == 0 {
				return ""
			}
			return strconv.Itoa(year)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"mod": func(a, b int) int {
			return a % b
		},
		"colorClass": func(color string) string {
			colors := map[string]string{
				"cyan":   "cyan-400",
				"pink":   "pink-400",
				"yellow": "yellow-400",
				"lime":   "lime-400",
				"orange": "orange-400",
				"purple": "purple-400",
				"red":    "red-400",
				"blue":   "blue-400",
				"green":  "green-400",
				"indigo": "indigo-400",
				"gray":   "gray-400",
				"black":  "gray-800",
			}
			if c, ok := colors[color]; ok {
				return c
			}
			return "gray-400"
		},
		"colorClassLight": func(color string) string {
			colors := map[string]string{
				"cyan":   "cyan-100",
				"pink":   "pink-100",
				"yellow": "yellow-100",
				"lime":   "lime-100",
				"orange": "orange-100",
				"purple": "purple-100",
				"red":    "red-100",
				"blue":   "blue-100",
				"green":  "green-100",
				"indigo": "indigo-100",
				"gray":   "gray-100",
				"black":  "gray-200",
			}
			if c, ok := colors[color]; ok {
				return c
			}
			return "gray-100"
		},
	}

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("views/*.html"))

	return &PortfolioHandler{
		service: svc,
		log:     log,
		tmpl:    tmpl,
	}
}

// RenderPortfolio renders the main portfolio page
func (h *PortfolioHandler) RenderPortfolio(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetPortfolioData(r.Context())
	if err != nil {
		h.log.Error("Failed to get portfolio data", zap.Error(err))
		http.Error(w, "Failed to load portfolio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		h.log.Error("Failed to render template", zap.Error(err))
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

// ================= Profile Handlers =================

// GetProfile returns the profile as JSON
func (h *PortfolioHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := h.service.GetProfile(r.Context())
	if err != nil {
		h.log.Error("Failed to get profile", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Profile not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Profile retrieved successfully", profile)
}

// CreateProfile creates a new profile
func (h *PortfolioHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req dto.ProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	profile, err := h.service.CreateProfile(r.Context(), &req)
	if err != nil {
		h.log.Error("Failed to create profile", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to create profile", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "Profile created successfully", profile)
}

// UpdateProfile updates the profile
func (h *PortfolioHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid profile ID", err.Error())
		return
	}

	var req dto.ProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	profile, err := h.service.UpdateProfile(r.Context(), id, &req)
	if err != nil {
		h.log.Error("Failed to update profile", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to update profile", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Profile updated successfully", profile)
}

// ================= Experience Handlers =================

// GetAllExperiences returns all experiences
func (h *PortfolioHandler) GetAllExperiences(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.service.GetAllExperiences(r.Context())
	if err != nil {
		h.log.Error("Failed to get experiences", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get experiences", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Experiences retrieved successfully", experiences)
}

// GetExperienceByID returns an experience by ID
func (h *PortfolioHandler) GetExperienceByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid experience ID", err.Error())
		return
	}

	exp, err := h.service.GetExperienceByID(r.Context(), id)
	if err != nil {
		h.log.Error("Failed to get experience", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Experience not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Experience retrieved successfully", exp)
}

// CreateExperience creates a new experience
func (h *PortfolioHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
	var req dto.ExperienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	exp, err := h.service.CreateExperience(r.Context(), &req)
	if err != nil {
		h.log.Error("Failed to create experience", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to create experience", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "Experience created successfully", exp)
}

// UpdateExperience updates an experience
func (h *PortfolioHandler) UpdateExperience(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid experience ID", err.Error())
		return
	}

	var req dto.ExperienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	exp, err := h.service.UpdateExperience(r.Context(), id, &req)
	if err != nil {
		h.log.Error("Failed to update experience", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to update experience", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Experience updated successfully", exp)
}

// DeleteExperience deletes an experience
func (h *PortfolioHandler) DeleteExperience(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid experience ID", err.Error())
		return
	}

	if err := h.service.DeleteExperience(r.Context(), id); err != nil {
		h.log.Error("Failed to delete experience", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to delete experience", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Experience deleted successfully", nil)
}

// ================= Skill Handlers =================

// GetAllSkills returns all skills
func (h *PortfolioHandler) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := h.service.GetAllSkills(r.Context())
	if err != nil {
		h.log.Error("Failed to get skills", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get skills", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Skills retrieved successfully", skills)
}

// GetSkillByID returns a skill by ID
func (h *PortfolioHandler) GetSkillByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid skill ID", err.Error())
		return
	}

	skill, err := h.service.GetSkillByID(r.Context(), id)
	if err != nil {
		h.log.Error("Failed to get skill", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Skill not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Skill retrieved successfully", skill)
}

// CreateSkill creates a new skill
func (h *PortfolioHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	var req dto.SkillRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	skill, err := h.service.CreateSkill(r.Context(), &req)
	if err != nil {
		h.log.Error("Failed to create skill", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to create skill", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "Skill created successfully", skill)
}

// UpdateSkill updates a skill
func (h *PortfolioHandler) UpdateSkill(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid skill ID", err.Error())
		return
	}

	var req dto.SkillRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	skill, err := h.service.UpdateSkill(r.Context(), id, &req)
	if err != nil {
		h.log.Error("Failed to update skill", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to update skill", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Skill updated successfully", skill)
}

// DeleteSkill deletes a skill
func (h *PortfolioHandler) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid skill ID", err.Error())
		return
	}

	if err := h.service.DeleteSkill(r.Context(), id); err != nil {
		h.log.Error("Failed to delete skill", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to delete skill", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Skill deleted successfully", nil)
}

// ================= Project Handlers =================

// GetAllProjects returns all projects
func (h *PortfolioHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		h.log.Error("Failed to get projects", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get projects", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Projects retrieved successfully", projects)
}

// GetProjectByID returns a project by ID
func (h *PortfolioHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	project, err := h.service.GetProjectByID(r.Context(), id)
	if err != nil {
		h.log.Error("Failed to get project", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Project not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Project retrieved successfully", project)
}

// CreateProject creates a new project
func (h *PortfolioHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req dto.ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	project, err := h.service.CreateProject(r.Context(), &req)
	if err != nil {
		h.log.Error("Failed to create project", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to create project", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "Project created successfully", project)
}

// UpdateProject updates a project
func (h *PortfolioHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	var req dto.ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	project, err := h.service.UpdateProject(r.Context(), id, &req)
	if err != nil {
		h.log.Error("Failed to update project", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to update project", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Project updated successfully", project)
}

// DeleteProject deletes a project
func (h *PortfolioHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid project ID", err.Error())
		return
	}

	if err := h.service.DeleteProject(r.Context(), id); err != nil {
		h.log.Error("Failed to delete project", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to delete project", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Project deleted successfully", nil)
}

// ================= Publication Handlers =================

// GetAllPublications returns all publications
func (h *PortfolioHandler) GetAllPublications(w http.ResponseWriter, r *http.Request) {
	pubs, err := h.service.GetAllPublications(r.Context())
	if err != nil {
		h.log.Error("Failed to get publications", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get publications", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Publications retrieved successfully", pubs)
}

// GetPublicationByID returns a publication by ID
func (h *PortfolioHandler) GetPublicationByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid publication ID", err.Error())
		return
	}

	pub, err := h.service.GetPublicationByID(r.Context(), id)
	if err != nil {
		h.log.Error("Failed to get publication", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Publication not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Publication retrieved successfully", pub)
}

// CreatePublication creates a new publication
func (h *PortfolioHandler) CreatePublication(w http.ResponseWriter, r *http.Request) {
	var req dto.PublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	pub, err := h.service.CreatePublication(r.Context(), &req)
	if err != nil {
		h.log.Error("Failed to create publication", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to create publication", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "Publication created successfully", pub)
}

// UpdatePublication updates a publication
func (h *PortfolioHandler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid publication ID", err.Error())
		return
	}

	var req dto.PublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	pub, err := h.service.UpdatePublication(r.Context(), id, &req)
	if err != nil {
		h.log.Error("Failed to update publication", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to update publication", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Publication updated successfully", pub)
}

// DeletePublication deletes a publication
func (h *PortfolioHandler) DeletePublication(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid publication ID", err.Error())
		return
	}

	if err := h.service.DeletePublication(r.Context(), id); err != nil {
		h.log.Error("Failed to delete publication", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to delete publication", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Publication deleted successfully", nil)
}

// ================= Contact Handler =================

// SubmitContact handles contact form submission
func (h *PortfolioHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	var req dto.ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Error("Failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.service.SubmitContact(r.Context(), &req); err != nil {
		h.log.Error("Failed to submit contact", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Failed to submit contact", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Contact submitted successfully", nil)
}

// ================= API Data Handler =================

// GetPortfolioData returns all portfolio data as JSON
func (h *PortfolioHandler) GetPortfolioData(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetPortfolioData(r.Context())
	if err != nil {
		h.log.Error("Failed to get portfolio data", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get portfolio data", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Portfolio data retrieved successfully", data)
}
