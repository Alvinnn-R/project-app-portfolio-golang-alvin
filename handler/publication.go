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

// PublicationHandler handles HTTP requests for publications
type PublicationHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewPublicationHandler creates a new publication handler
func NewPublicationHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *PublicationHandler {
	return &PublicationHandler{
		service: svc,
		log:     log,
	}
}

// GetAllPublications returns all publications
func (h *PublicationHandler) GetAllPublications(w http.ResponseWriter, r *http.Request) {
	pubs, err := h.service.GetAllPublications(r.Context())
	if err != nil {
		h.log.Error("Failed to get publications", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to get publications", err.Error())
		return
	}
	utils.ResponseSuccess(w, http.StatusOK, "Publications retrieved successfully", pubs)
}

// GetPublicationByID returns a publication by ID
func (h *PublicationHandler) GetPublicationByID(w http.ResponseWriter, r *http.Request) {
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
func (h *PublicationHandler) CreatePublication(w http.ResponseWriter, r *http.Request) {
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
func (h *PublicationHandler) UpdatePublication(w http.ResponseWriter, r *http.Request) {
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
func (h *PublicationHandler) DeletePublication(w http.ResponseWriter, r *http.Request) {
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
