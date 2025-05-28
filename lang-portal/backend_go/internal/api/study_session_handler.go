package api

import (
	"net/http"
	"strconv"

	"github.com/erans/lang-portal/internal/models"
	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// StudySessionHandler handles study session-related HTTP requests
type StudySessionHandler struct {
	sessionService *service.StudySessionService
}

// NewStudySessionHandler creates a new StudySessionHandler
func NewStudySessionHandler(sessionService *service.StudySessionService) *StudySessionHandler {
	return &StudySessionHandler{sessionService: sessionService}
}

// RegisterRoutes registers the study session routes
func (h *StudySessionHandler) RegisterRoutes(router *gin.Engine) {
	sessions := router.Group("/api/sessions")
	{
		sessions.GET("", h.ListSessions)
		sessions.GET("/:id", h.GetSession)
		sessions.POST("", h.CreateSession)
		sessions.PUT("/:id", h.UpdateSession)
		sessions.PUT("/:id/end", h.EndSession)
		sessions.GET("/:id/review-items", h.GetSessionReviewItems)
	}
}

// ListSessions handles GET /api/sessions
func (h *StudySessionHandler) ListSessions(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	sessions, err := h.sessionService.ListSessions(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

// GetSession handles GET /api/sessions/:id
func (h *StudySessionHandler) GetSession(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	session, err := h.sessionService.GetSession(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

// CreateSession handles POST /api/sessions
func (h *StudySessionHandler) CreateSession(c *gin.Context) {
	var session models.StudySession
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sessionService.CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}

// UpdateSession handles PUT /api/sessions/:id
func (h *StudySessionHandler) UpdateSession(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	var session models.StudySession
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session.ID = id
	if err := h.sessionService.UpdateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

// EndSession handles PUT /api/sessions/:id/end
func (h *StudySessionHandler) EndSession(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	var request struct {
		Score float64 `json:"score" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sessionService.EndSession(id, request.Score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetSessionReviewItems handles GET /api/sessions/:id/review-items
func (h *StudySessionHandler) GetSessionReviewItems(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	items, err := h.sessionService.GetSessionReviewItems(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
