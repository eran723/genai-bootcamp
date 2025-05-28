package api

import (
	"net/http"

	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// DashboardHandler handles HTTP requests for dashboard data
type DashboardHandler struct {
	dashboardService *service.DashboardService
}

// NewDashboardHandler creates a new DashboardHandler
func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService}
}

// RegisterRoutes registers the dashboard routes
func (h *DashboardHandler) RegisterRoutes(r *gin.Engine) {
	dashboard := r.Group("/api/dashboard")
	{
		dashboard.GET("/last_session", h.GetLastSession)
		dashboard.GET("/stats", h.GetStats)
		dashboard.GET("/progress", h.GetProgress)
	}
}

// GetLastSession handles GET /api/dashboard/last_session
func (h *DashboardHandler) GetLastSession(c *gin.Context) {
	session, err := h.dashboardService.GetLastSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if session == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No study sessions found",
		})
		return
	}

	c.JSON(http.StatusOK, session)
}

// GetStats handles GET /api/dashboard/stats
func (h *DashboardHandler) GetStats(c *gin.Context) {
	stats, err := h.dashboardService.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetProgress handles GET /api/dashboard/progress
func (h *DashboardHandler) GetProgress(c *gin.Context) {
	progress, err := h.dashboardService.GetProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}
