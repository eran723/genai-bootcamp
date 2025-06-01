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
	sessions := router.Group("/api/study-sessions")
	{
		sessions.GET("", h.ListSessions)
		sessions.GET("/:id", h.GetSession)
		sessions.POST("", h.CreateSession)
		sessions.PUT("/:id", h.UpdateSession)
		sessions.PUT("/:id/end", h.EndSession)
		sessions.GET("/:id/words", h.GetSessionWords)
		sessions.GET("/:id/review-items", h.GetSessionReviewItems)
	}
}

// ListSessions handles GET /api/study-sessions
func (h *StudySessionHandler) ListSessions(c *gin.Context) {
	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	// Get sessions from service
	result, err := h.sessionService.ListSessions(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate pagination metadata
	totalPages := int((result.TotalItems + int64(limit) - 1) / int64(limit))

	response := models.PaginatedResponse{
		Items: result.Items,
		Pagination: models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   result.TotalItems,
			ItemsPerPage: limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetSession handles GET /api/study-sessions/:id
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

// GetSessionWords handles GET /api/study-sessions/:id/words
func (h *StudySessionHandler) GetSessionWords(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	items, err := h.sessionService.GetSessionReviewItems(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages based on total items
	totalItems := int64(len(items))
	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	// Paginate the results
	start := offset
	end := offset + limit
	if start > len(items) {
		start = len(items)
	}
	if end > len(items) {
		end = len(items)
	}

	response := models.PaginatedResponse{
		Items: items[start:end],
		Pagination: models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetSessionReviewItems handles GET /api/study-sessions/:id/review-items
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

// CreateSession handles POST /api/study-sessions
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

// UpdateSession handles PUT /api/study-sessions/:id
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

// EndSession handles PUT /api/study-sessions/:id/end
func (h *StudySessionHandler) EndSession(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	var payload struct {
		Score float64 `json:"score" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sessionService.EndSession(id, payload.Score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
