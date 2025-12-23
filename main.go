package main

import (
	"fmt"
	"log"
	"net/http"
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

	// Initialize layers
	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, logger)

	// Create router
	r := router.NewRouter(h, svc, logger)

	// Start server
	fmt.Println("=================================")
	fmt.Println("Portfolio Server is running...")
	fmt.Println("Open http://localhost:8080 in your browser")
	fmt.Println("API available at http://localhost:8080/api/v1")
	fmt.Println("=================================")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error:", err)
	}
}
