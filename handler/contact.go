package handler

import (
	"encoding/json"
	"net/http"
	"session-19/dto"
	"session-19/service"
	"session-19/utils"

	"go.uber.org/zap"
)

// ContactHandler handles HTTP requests for contact form
type ContactHandler struct {
	service service.PortfolioServiceInterface
	log     *zap.Logger
}

// NewContactHandler creates a new contact handler
func NewContactHandler(svc service.PortfolioServiceInterface, log *zap.Logger) *ContactHandler {
	return &ContactHandler{
		service: svc,
		log:     log,
	}
}

// SubmitContact handles contact form submission
func (h *ContactHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
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
