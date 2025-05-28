package main

import (
	"log"

	"github.com/erans/lang-portal/internal/api"
	"github.com/erans/lang-portal/internal/database"
	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := database.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// Run migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Create services
	wordService := service.NewWordService(database.GetDB())
	groupService := service.NewGroupService(database.GetDB())
	dashboardService := service.NewDashboardService(database.GetDB())
	studySessionService := service.NewStudySessionService(database.GetDB())
	studyActivityService := service.NewStudyActivityService(database.GetDB())
	systemService := service.NewSystemService(database.GetDB())

	// Create handlers
	wordHandler := api.NewWordHandler(wordService)
	groupHandler := api.NewGroupHandler(groupService)
	dashboardHandler := api.NewDashboardHandler(dashboardService)
	studySessionHandler := api.NewStudySessionHandler(studySessionService)
	studyActivityHandler := api.NewStudyActivityHandler(studyActivityService)
	systemHandler := api.NewSystemHandler(systemService)

	// Create default gin engine with middleware
	r := gin.Default()

	// Add CORS middleware
	r.Use(corsMiddleware())

	// Register all routes
	wordHandler.RegisterRoutes(r)
	groupHandler.RegisterRoutes(r)
	dashboardHandler.RegisterRoutes(r)
	studySessionHandler.RegisterRoutes(r)
	studyActivityHandler.RegisterRoutes(r)
	systemHandler.RegisterRoutes(r)

	// Add basic health check
	r.GET("/health", healthCheck)

	log.Println("Server starting on :8080...")
	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// corsMiddleware handles CORS headers
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// healthCheck handles the health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"version": "1.0.0",
	})
}

func setupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API group
	api := r.Group("/api")
	{
		// Dashboard endpoints
		api.GET("/dashboard/last_session", func(c *gin.Context) {
			// TODO: Implement last session endpoint
			c.JSON(501, gin.H{"error": "Not implemented"})
		})

		api.GET("/dashboard/stats", func(c *gin.Context) {
			// TODO: Implement stats endpoint
			c.JSON(501, gin.H{"error": "Not implemented"})
		})

		api.GET("/dashboard/progress", func(c *gin.Context) {
			// TODO: Implement progress endpoint
			c.JSON(501, gin.H{"error": "Not implemented"})
		})
	}
}
