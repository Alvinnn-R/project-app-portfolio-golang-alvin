package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"session-19/database"
	"session-19/handler"
	"session-19/repository"
	"session-19/router"
	"session-19/service"
	"session-19/utils"
)

func main() {
	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close(nil)

	// Initialize logger
	logger, err := utils.InitLogger("./logs/app-", true)
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Sync()

	// Parse templates from layouts and pages directories
	tmpl, err := parseTemplates()
	if err != nil {
		log.Fatal("Failed to parse templates:", err)
	}

	// Initialize layers
	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, logger, tmpl)

	// Create router
	r := router.NewRouter(h, svc, logger)

	// Start server
	fmt.Println("=================================")
	fmt.Println("Portfolio Server is running...")
	fmt.Println("Open http://localhost:8080 in your browser")
	fmt.Println("Admin Panel: http://localhost:8080/login")
	fmt.Println("API available at http://localhost:8080/api/v1")
	fmt.Println("=================================")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error:", err)
	}
}

// parseTemplates parses all templates from layouts and pages directories
func parseTemplates() (*template.Template, error) {
	tmpl := template.New("")

	// Parse layouts first
	layouts, err := filepath.Glob("views/layouts/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to glob layouts: %w", err)
	}

	// Parse pages
	pages, err := filepath.Glob("views/pages/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to glob pages: %w", err)
	}

	// Combine all template files
	allFiles := append(layouts, pages...)

	// Parse all templates
	tmpl, err = tmpl.ParseFiles(allFiles...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return tmpl, nil
}
