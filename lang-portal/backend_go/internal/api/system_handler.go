package api

import (
	"net/http"

	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// SystemHandler handles system-wide HTTP requests
type SystemHandler struct {
	systemService *service.SystemService
}

// NewSystemHandler creates a new SystemHandler
func NewSystemHandler(systemService *service.SystemService) *SystemHandler {
	return &SystemHandler{systemService: systemService}
}

// RegisterRoutes registers the system routes
func (h *SystemHandler) RegisterRoutes(router *gin.Engine) {
	system := router.Group("/api/system")
	{
		system.GET("/stats", h.GetSystemStats)
		system.GET("/health", h.GetSystemHealth)
		system.POST("/backup", h.BackupDatabase)
		system.GET("/database/size", h.GetDatabaseSize)
		system.GET("/backup/last", h.GetLastBackupInfo)
		system.POST("/prune", h.PruneOldData)
	}
}

// GetSystemStats handles GET /api/system/stats
func (h *SystemHandler) GetSystemStats(c *gin.Context) {
	stats, err := h.systemService.GetSystemStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetSystemHealth handles GET /api/system/health
func (h *SystemHandler) GetSystemHealth(c *gin.Context) {
	health, err := h.systemService.GetSystemHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if health.Status != "healthy" {
		c.JSON(http.StatusServiceUnavailable, health)
		return
	}

	c.JSON(http.StatusOK, health)
}

// BackupDatabase handles POST /api/system/backup
func (h *SystemHandler) BackupDatabase(c *gin.Context) {
	var request struct {
		BackupPath string `json:"backup_path" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.systemService.BackupDatabase(request.BackupPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetDatabaseSize handles GET /api/system/database/size
func (h *SystemHandler) GetDatabaseSize(c *gin.Context) {
	size, err := h.systemService.GetDatabaseSize()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"size_bytes": size})
}

// GetLastBackupInfo handles GET /api/system/backup/last
func (h *SystemHandler) GetLastBackupInfo(c *gin.Context) {
	info, err := h.systemService.GetLastBackupInfo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, info)
}

// PruneOldData handles POST /api/system/prune
func (h *SystemHandler) PruneOldData(c *gin.Context) {
	var request struct {
		RetentionDays int `json:"retention_days" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.systemService.PruneOldData(request.RetentionDays); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
