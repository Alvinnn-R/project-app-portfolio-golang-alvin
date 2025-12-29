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

// SkillHandler handles HTTP requests for skills
type SkillHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewSkillHandler creates a new skill handler
func NewSkillHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *SkillHandler {
	return &SkillHandler{
		service: svc,
		log:     log,
	}
}

// GetAllSkills returns all skills
func (h *SkillHandler) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := h.service.GetAllSkills(r.Context())
	if err != nil {
		h.log.Error("Failed to get skills", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get skills", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Skills retrieved successfully", skills)
}

// GetSkillByID returns a skill by ID
func (h *SkillHandler) GetSkillByID(w http.ResponseWriter, r *http.Request) {
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
func (h *SkillHandler) CreateSkill(w http.ResponseWriter, r *http.Request) {
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
func (h *SkillHandler) UpdateSkill(w http.ResponseWriter, r *http.Request) {
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
func (h *SkillHandler) DeleteSkill(w http.ResponseWriter, r *http.Request) {
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
