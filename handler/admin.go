package handler

import (
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

// AdminHandler handles admin panel requests
type AdminHandler struct {
	portfolioService service.PortfolioServiceInterface
	log              *zap.Logger
	tmpl             *template.Template
}

// NewAdminHandler creates a new admin handler
func NewAdminHandler(portfolioService service.PortfolioServiceInterface, log *zap.Logger, tmpl *template.Template) *AdminHandler {
	return &AdminHandler{
		portfolioService: portfolioService,
		log:              log,
		tmpl:             tmpl,
	}
}

// Dashboard renders the admin dashboard
func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.portfolioService.GetPortfolioData(ctx)
	if err != nil {
		h.log.Error("Failed to get portfolio data", zap.Error(err))
	}

	// Count items
	stats := map[string]int{
		"experiences":  0,
		"skills":       0,
		"projects":     0,
		"publications": 0,
	}

	if data != nil {
		stats["experiences"] = len(data.Experiences)
		for _, skills := range data.Skills {
			stats["skills"] += len(skills)
		}
		stats["projects"] = len(data.Projects)
		stats["publications"] = len(data.Publications)
	}

	if err := h.tmpl.ExecuteTemplate(w, "dashboard", map[string]interface{}{
		"Stats":   stats,
		"Profile": data.Profile,
	}); err != nil {
		h.log.Error("Failed to render dashboard", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ==================== PROFILE ====================

// ProfileEdit renders the profile edit page
func (h *AdminHandler) ProfileEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile, err := h.portfolioService.GetProfile(ctx)
	if err != nil {
		h.log.Warn("No profile found", zap.Error(err))
	}

	if err := h.tmpl.ExecuteTemplate(w, "profile_form", map[string]interface{}{
		"Profile": profile,
	}); err != nil {
		h.log.Error("Failed to render profile form", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ProfileSave handles profile create/update
func (h *AdminHandler) ProfileSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/profile", http.StatusSeeOther)
		return
	}

	// Parse multipart form (max 10MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.log.Error("Failed to parse form", zap.Error(err))
	}

	ctx := r.Context()

	// Handle photo upload
	photoURL := r.FormValue("existing_photo")
	if file, header, err := r.FormFile("photo"); err == nil {
		defer file.Close()
		uploadedPath, uploadErr := utils.UploadFile(file, header, "uploads/profile")
		if uploadErr != nil {
			h.log.Error("Failed to upload photo", zap.Error(uploadErr))
			h.renderProfileError(w, &dto.ProfileRequest{
				Name:        r.FormValue("name"),
				Title:       r.FormValue("title"),
				Description: r.FormValue("description"),
				PhotoURL:    photoURL,
				Email:       r.FormValue("email"),
				LinkedInURL: r.FormValue("linkedin_url"),
				GithubURL:   r.FormValue("github_url"),
				CVURL:       r.FormValue("cv_url"),
			}, uploadErr.Error())
			return
		}
		photoURL = uploadedPath
	}

	req := &dto.ProfileRequest{
		Name:        r.FormValue("name"),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		PhotoURL:    photoURL,
		Email:       r.FormValue("email"),
		LinkedInURL: r.FormValue("linkedin_url"),
		GithubURL:   r.FormValue("github_url"),
		CVURL:       r.FormValue("cv_url"),
	}

	idStr := r.FormValue("id")
	if idStr != "" && idStr != "0" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		_, err := h.portfolioService.UpdateProfile(ctx, id, req)
		if err != nil {
			h.renderProfileError(w, req, err.Error())
			return
		}
	} else {
		_, err := h.portfolioService.CreateProfile(ctx, req)
		if err != nil {
			h.renderProfileError(w, req, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/admin/dashboard?success=profile", http.StatusSeeOther)
}

func (h *AdminHandler) renderProfileError(w http.ResponseWriter, req *dto.ProfileRequest, errMsg string) {
	h.tmpl.ExecuteTemplate(w, "profile_form", map[string]interface{}{
		"Error":   errMsg,
		"Profile": req,
	})
}

// ==================== EXPERIENCES ====================

// ExperiencesList renders the experiences list
func (h *AdminHandler) ExperiencesList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	experiences, err := h.portfolioService.GetAllExperiences(ctx)
	if err != nil {
		h.log.Error("Failed to get experiences", zap.Error(err))
	}

	if err := h.tmpl.ExecuteTemplate(w, "experiences_list", map[string]interface{}{
		"Experiences": experiences,
		"Success":     r.URL.Query().Get("success"),
	}); err != nil {
		h.log.Error("Failed to render experiences list", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ExperienceForm renders the experience form
func (h *AdminHandler) ExperienceForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	var experience interface{}
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		exp, err := h.portfolioService.GetExperienceByID(ctx, id)
		if err == nil {
			experience = exp
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, "experience_form", map[string]interface{}{
		"Experience": experience,
		"Types":      []string{"work", "internship", "campus", "competition"},
	}); err != nil {
		h.log.Error("Failed to render experience form", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ExperienceSave handles experience create/update
func (h *AdminHandler) ExperienceSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/experiences", http.StatusSeeOther)
		return
	}

	ctx := r.Context()
	req := &dto.ExperienceRequest{
		Title:        r.FormValue("title"),
		Organization: r.FormValue("organization"),
		Period:       r.FormValue("period"),
		Description:  r.FormValue("description"),
		Type:         r.FormValue("type"),
		Color:        r.FormValue("color"),
	}

	idStr := r.FormValue("id")
	if idStr != "" && idStr != "0" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		_, err := h.portfolioService.UpdateExperience(ctx, id, req)
		if err != nil {
			h.renderExperienceError(w, req, err.Error(), nil)
			return
		}
	} else {
		_, err := h.portfolioService.CreateExperience(ctx, req)
		if err != nil {
			h.renderExperienceError(w, req, err.Error(), nil)
			return
		}
	}

	http.Redirect(w, r, "/admin/experiences?success=saved", http.StatusSeeOther)
}

// ExperienceDelete handles experience deletion
func (h *AdminHandler) ExperienceDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.portfolioService.DeleteExperience(ctx, id); err != nil {
		h.log.Error("Failed to delete experience", zap.Error(err))
	}

	http.Redirect(w, r, "/admin/experiences?success=deleted", http.StatusSeeOther)
}

func (h *AdminHandler) renderExperienceError(w http.ResponseWriter, req *dto.ExperienceRequest, errMsg string, exp interface{}) {
	h.tmpl.ExecuteTemplate(w, "experience_form", map[string]interface{}{
		"Error":      errMsg,
		"Experience": req,
		"Types":      []string{"work", "internship", "campus", "competition"},
	})
}

// ==================== SKILLS ====================

// SkillsList renders the skills list
func (h *AdminHandler) SkillsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	skills, err := h.portfolioService.GetAllSkills(ctx)
	if err != nil {
		h.log.Error("Failed to get skills", zap.Error(err))
	}

	// Group by category
	grouped := make(map[string]interface{})
	for _, skill := range skills {
		cat := skill.Category
		if _, ok := grouped[cat]; !ok {
			grouped[cat] = []interface{}{}
		}
		grouped[cat] = append(grouped[cat].([]interface{}), skill)
	}

	if err := h.tmpl.ExecuteTemplate(w, "skills_list", map[string]interface{}{
		"Skills":  skills,
		"Grouped": grouped,
		"Success": r.URL.Query().Get("success"),
	}); err != nil {
		h.log.Error("Failed to render skills list", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// SkillForm renders the skill form
func (h *AdminHandler) SkillForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	var skill interface{}
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		s, err := h.portfolioService.GetSkillByID(ctx, id)
		if err == nil {
			skill = s
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, "skill_form", map[string]interface{}{
		"Skill":  skill,
		"Levels": []string{"beginner", "intermediate", "advanced"},
	}); err != nil {
		h.log.Error("Failed to render skill form", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// SkillSave handles skill create/update
func (h *AdminHandler) SkillSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/skills", http.StatusSeeOther)
		return
	}

	ctx := r.Context()
	req := &dto.SkillRequest{
		Category: r.FormValue("category"),
		Name:     r.FormValue("name"),
		Level:    r.FormValue("level"),
		Color:    r.FormValue("color"),
	}

	idStr := r.FormValue("id")
	if idStr != "" && idStr != "0" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		_, err := h.portfolioService.UpdateSkill(ctx, id, req)
		if err != nil {
			h.renderSkillError(w, req, err.Error())
			return
		}
	} else {
		_, err := h.portfolioService.CreateSkill(ctx, req)
		if err != nil {
			h.renderSkillError(w, req, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/admin/skills?success=saved", http.StatusSeeOther)
}

// SkillDelete handles skill deletion
func (h *AdminHandler) SkillDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.portfolioService.DeleteSkill(ctx, id); err != nil {
		h.log.Error("Failed to delete skill", zap.Error(err))
	}

	http.Redirect(w, r, "/admin/skills?success=deleted", http.StatusSeeOther)
}

func (h *AdminHandler) renderSkillError(w http.ResponseWriter, req *dto.SkillRequest, errMsg string) {
	h.tmpl.ExecuteTemplate(w, "skill_form", map[string]interface{}{
		"Error":  errMsg,
		"Skill":  req,
		"Levels": []string{"beginner", "intermediate", "advanced"},
	})
}

// ==================== PROJECTS ====================

// ProjectsList renders the projects list
func (h *AdminHandler) ProjectsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projects, err := h.portfolioService.GetAllProjects(ctx)
	if err != nil {
		h.log.Error("Failed to get projects", zap.Error(err))
	}

	if err := h.tmpl.ExecuteTemplate(w, "projects_list", map[string]interface{}{
		"Projects": projects,
		"Success":  r.URL.Query().Get("success"),
	}); err != nil {
		h.log.Error("Failed to render projects list", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ProjectForm renders the project form
func (h *AdminHandler) ProjectForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	var project interface{}
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		p, err := h.portfolioService.GetProjectByID(ctx, id)
		if err == nil {
			project = p
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, "project_form", map[string]interface{}{
		"Project": project,
	}); err != nil {
		h.log.Error("Failed to render project form", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ProjectSave handles project create/update
func (h *AdminHandler) ProjectSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/projects", http.StatusSeeOther)
		return
	}

	// Parse multipart form (max 10MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.log.Error("Failed to parse form", zap.Error(err))
	}

	ctx := r.Context()

	// Handle image upload
	imageURL := r.FormValue("existing_image")
	if file, header, err := r.FormFile("image"); err == nil {
		defer file.Close()
		uploadedPath, uploadErr := utils.UploadFile(file, header, "uploads/projects")
		if uploadErr != nil {
			h.log.Error("Failed to upload project image", zap.Error(uploadErr))
			h.renderProjectError(w, &dto.ProjectRequest{
				Title:       r.FormValue("title"),
				Description: r.FormValue("description"),
				ImageURL:    imageURL,
				ProjectURL:  r.FormValue("project_url"),
				GithubURL:   r.FormValue("github_url"),
				TechStack:   r.FormValue("tech_stack"),
				Color:       r.FormValue("color"),
			}, uploadErr.Error())
			return
		}
		imageURL = uploadedPath
	}

	req := &dto.ProjectRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		ImageURL:    imageURL,
		ProjectURL:  r.FormValue("project_url"),
		GithubURL:   r.FormValue("github_url"),
		TechStack:   r.FormValue("tech_stack"),
		Color:       r.FormValue("color"),
	}

	// Get profile ID for foreign key
	profile, err := h.portfolioService.GetProfile(ctx)
	if err == nil && profile != nil {
		req.ProfileID = profile.ID
	}

	idStr := r.FormValue("id")
	if idStr != "" && idStr != "0" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		_, err := h.portfolioService.UpdateProject(ctx, id, req)
		if err != nil {
			h.renderProjectError(w, req, err.Error())
			return
		}
	} else {
		_, err := h.portfolioService.CreateProject(ctx, req)
		if err != nil {
			h.renderProjectError(w, req, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/admin/projects?success=saved", http.StatusSeeOther)
}

// ProjectDelete handles project deletion
func (h *AdminHandler) ProjectDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.portfolioService.DeleteProject(ctx, id); err != nil {
		h.log.Error("Failed to delete project", zap.Error(err))
	}

	http.Redirect(w, r, "/admin/projects?success=deleted", http.StatusSeeOther)
}

func (h *AdminHandler) renderProjectError(w http.ResponseWriter, req *dto.ProjectRequest, errMsg string) {
	h.tmpl.ExecuteTemplate(w, "project_form", map[string]interface{}{
		"Error":   errMsg,
		"Project": req,
	})
}

// ==================== PUBLICATIONS ====================

// PublicationsList renders the publications list
func (h *AdminHandler) PublicationsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	publications, err := h.portfolioService.GetAllPublications(ctx)
	if err != nil {
		h.log.Error("Failed to get publications", zap.Error(err))
	}

	if err := h.tmpl.ExecuteTemplate(w, "publications_list", map[string]interface{}{
		"Publications": publications,
		"Success":      r.URL.Query().Get("success"),
	}); err != nil {
		h.log.Error("Failed to render publications list", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// PublicationForm renders the publication form
func (h *AdminHandler) PublicationForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")

	var publication interface{}
	if idStr != "" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		p, err := h.portfolioService.GetPublicationByID(ctx, id)
		if err == nil {
			publication = p
		}
	}

	if err := h.tmpl.ExecuteTemplate(w, "publication_form", map[string]interface{}{
		"Publication": publication,
	}); err != nil {
		h.log.Error("Failed to render publication form", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// PublicationSave handles publication create/update
func (h *AdminHandler) PublicationSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/admin/publications", http.StatusSeeOther)
		return
	}

	ctx := r.Context()
	year, _ := strconv.Atoi(r.FormValue("year"))
	req := &dto.PublicationRequest{
		Title:          r.FormValue("title"),
		Authors:        r.FormValue("authors"),
		Journal:        r.FormValue("journal"),
		Year:           year,
		Description:    r.FormValue("description"),
		ImageURL:       r.FormValue("image_url"),
		PublicationURL: r.FormValue("publication_url"),
		Color:          r.FormValue("color"),
	}

	idStr := r.FormValue("id")
	if idStr != "" && idStr != "0" {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		_, err := h.portfolioService.UpdatePublication(ctx, id, req)
		if err != nil {
			h.renderPublicationError(w, req, err.Error())
			return
		}
	} else {
		_, err := h.portfolioService.CreatePublication(ctx, req)
		if err != nil {
			h.renderPublicationError(w, req, err.Error())
			return
		}
	}

	http.Redirect(w, r, "/admin/publications?success=saved", http.StatusSeeOther)
}

// PublicationDelete handles publication deletion
func (h *AdminHandler) PublicationDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := h.portfolioService.DeletePublication(ctx, id); err != nil {
		h.log.Error("Failed to delete publication", zap.Error(err))
	}

	http.Redirect(w, r, "/admin/publications?success=deleted", http.StatusSeeOther)
}

func (h *AdminHandler) renderPublicationError(w http.ResponseWriter, req *dto.PublicationRequest, errMsg string) {
	h.tmpl.ExecuteTemplate(w, "publication_form", map[string]interface{}{
		"Error":       errMsg,
		"Publication": req,
	})
}

// Helper function to get template FuncMap for admin templates
func GetAdminTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"title": strings.Title,
	}
}
