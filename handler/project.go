package handler

import (
	"encoding/json"
	"net/http"
	"session-19/dto"
	"session-19/service"
	"session-19/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// ProjectHandler handles HTTP requests for projects
type ProjectHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *ProjectHandler {
	return &ProjectHandler{
		service: svc,
		log:     log,
	}
}

// GetAllProjects returns all projects
func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.service.GetAllProjects(r.Context())
	if err != nil {
		h.log.Error("Failed to get projects", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get projects", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Projects retrieved successfully", projects)
}

// GetProjectByID returns a project by ID
func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
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
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
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
func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
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
func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
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
