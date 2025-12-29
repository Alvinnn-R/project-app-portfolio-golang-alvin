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

// ProfileHandler handles HTTP requests for profile
type ProfileHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewProfileHandler creates a new profile handler
func NewProfileHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *ProfileHandler {
	return &ProfileHandler{
		service: svc,
		log:     log,
	}
}

// GetProfile returns the profile as JSON
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := h.service.GetProfile(r.Context())
	if err != nil {
		h.log.Error("Failed to get profile", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "Profile not found", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Profile retrieved successfully", profile)
}

// CreateProfile creates a new profile
func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
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
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
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
