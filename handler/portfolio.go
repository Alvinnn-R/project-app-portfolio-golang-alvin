package handler

import (
	"html/template"
	"net/http"
	"session-19/service"
	"session-19/utils"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

// PortfolioHandler handles HTTP requests for portfolio main page and data
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
