package handler

import (
	"html/template"
	"net/http"
	"session-19/dto"
	"session-19/service"
	"strconv"

	"go.uber.org/zap"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService service.AuthServiceInterface
	log         *zap.Logger
	tmpl        *template.Template
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService service.AuthServiceInterface, log *zap.Logger, tmpl *template.Template) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         log,
		tmpl:        tmpl,
	}
}

// LoginView renders the login page
func (h *AuthHandler) LoginView(w http.ResponseWriter, r *http.Request) {
	// Check if already logged in
	if c, err := r.Cookie("session"); err == nil && c.Value != "" {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	if err := h.tmpl.ExecuteTemplate(w, "login", nil); err != nil {
		h.log.Error("Failed to render login page", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Login handles login form submission
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	req := &dto.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	user, err := h.authService.Login(r.Context(), req)
	if err != nil {
		h.log.Warn("Login failed", zap.String("email", req.Email), zap.Error(err))
		h.tmpl.ExecuteTemplate(w, "login", map[string]interface{}{
			"Error": err.Error(),
			"Email": req.Email,
		})
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "portfolio-" + strconv.FormatInt(user.ID, 10),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400 * 7, // 7 days
	})

	h.log.Info("User logged in", zap.String("email", user.Email))
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

// LogoutView renders the logout confirmation page
func (h *AuthHandler) LogoutView(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "logout", nil); err != nil {
		h.log.Error("Failed to render logout page", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Logout handles logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	h.log.Info("User logged out")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Page401 renders the unauthorized page
func (h *AuthHandler) Page401(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	if err := h.tmpl.ExecuteTemplate(w, "page401", nil); err != nil {
		h.log.Error("Failed to render 401 page", zap.Error(err))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
