package api

import (
	"net/http"
	"strconv"

	"github.com/erans/lang-portal/internal/models"
	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// StudyActivityHandler handles study activity-related HTTP requests
type StudyActivityHandler struct {
	activityService *service.StudyActivityService
}

// NewStudyActivityHandler creates a new StudyActivityHandler
func NewStudyActivityHandler(activityService *service.StudyActivityService) *StudyActivityHandler {
	return &StudyActivityHandler{activityService: activityService}
}

// RegisterRoutes registers the study activity routes
func (h *StudyActivityHandler) RegisterRoutes(router *gin.Engine) {
	activities := router.Group("/api/activities")
	{
		activities.GET("", h.ListActivities)
		activities.GET("/:id", h.GetActivity)
		activities.POST("", h.CreateActivity)
		activities.PUT("/:id", h.UpdateActivity)
		activities.DELETE("/:id", h.DeleteActivity)
		activities.GET("/:id/sessions", h.GetActivitySessions)
	}
}

// ListActivities handles GET /api/activities
func (h *StudyActivityHandler) ListActivities(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	activities, err := h.activityService.ListActivities(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

// GetActivity handles GET /api/activities/:id
func (h *StudyActivityHandler) GetActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	activity, err := h.activityService.GetActivity(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// CreateActivity handles POST /api/activities
func (h *StudyActivityHandler) CreateActivity(c *gin.Context) {
	var activity models.StudyActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.activityService.CreateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

// UpdateActivity handles PUT /api/activities/:id
func (h *StudyActivityHandler) UpdateActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	var activity models.StudyActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity.ID = id
	if err := h.activityService.UpdateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// DeleteActivity handles DELETE /api/activities/:id
func (h *StudyActivityHandler) DeleteActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	if err := h.activityService.DeleteActivity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetActivitySessions handles GET /api/activities/:id/sessions
func (h *StudyActivityHandler) GetActivitySessions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	sessions, err := h.activityService.GetActivitySessions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessions)
}
