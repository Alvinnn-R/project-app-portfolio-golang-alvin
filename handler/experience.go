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

// ExperienceHandler handles HTTP requests for experiences
type ExperienceHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewExperienceHandler creates a new experience handler
func NewExperienceHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *ExperienceHandler {
	return &ExperienceHandler{
		service: svc,
		log:     log,
	}
}

// GetAllExperiences returns all experiences
func (h *ExperienceHandler) GetAllExperiences(w http.ResponseWriter, r *http.Request) {
	experiences, err := h.service.GetAllExperiences(r.Context())
	if err != nil {
		h.log.Error("Failed to get experiences", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get experiences", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Experiences retrieved successfully", experiences)
}

// GetExperienceByID returns an experience by ID
func (h *ExperienceHandler) GetExperienceByID(w http.ResponseWriter, r *http.Request) {
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
func (h *ExperienceHandler) CreateExperience(w http.ResponseWriter, r *http.Request) {
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
func (h *ExperienceHandler) UpdateExperience(w http.ResponseWriter, r *http.Request) {
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
func (h *ExperienceHandler) DeleteExperience(w http.ResponseWriter, r *http.Request) {
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
