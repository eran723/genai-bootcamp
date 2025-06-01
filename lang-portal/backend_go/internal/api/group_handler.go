package api

import (
	"net/http"
	"strconv"

	"github.com/erans/lang-portal/internal/models"
	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// GroupHandler handles group-related HTTP requests
type GroupHandler struct {
	groupService *service.GroupService
}

// NewGroupHandler creates a new GroupHandler
func NewGroupHandler(groupService *service.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

// RegisterRoutes registers the group routes
func (h *GroupHandler) RegisterRoutes(router *gin.Engine) {
	groups := router.Group("/api/groups")
	{
		groups.GET("", h.ListGroups)
		groups.GET("/:id", h.GetGroup)
		groups.POST("", h.CreateGroup)
		groups.PUT("/:id", h.UpdateGroup)
		groups.DELETE("/:id", h.DeleteGroup)
		groups.GET("/:id/words", h.GetGroupWords)
		groups.GET("/:id/study-sessions", h.GetGroupStudySessions)
	}
}

// ListGroups handles GET /api/groups
func (h *GroupHandler) ListGroups(c *gin.Context) {
	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	// Get groups from service
	result, err := h.groupService.ListGroups(offset, limit)
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

// GetGroup handles GET /api/groups/:id
func (h *GroupHandler) GetGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	group, err := h.groupService.GetGroup(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetGroupWords handles GET /api/groups/:id/words
func (h *GroupHandler) GetGroupWords(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	words, err := h.groupService.GetGroupWords(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages based on total words
	totalItems := int64(len(words))
	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	// Paginate the results
	start := offset
	end := offset + limit
	if start > len(words) {
		start = len(words)
	}
	if end > len(words) {
		end = len(words)
	}

	response := models.PaginatedResponse{
		Items: words[start:end],
		Pagination: models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetGroupStudySessions handles GET /api/groups/:id/study-sessions
func (h *GroupHandler) GetGroupStudySessions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	sessions, err := h.groupService.GetGroupStudySessions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages based on total sessions
	totalItems := int64(len(sessions))
	totalPages := int((totalItems + int64(limit) - 1) / int64(limit))

	// Paginate the results
	start := offset
	end := offset + limit
	if start > len(sessions) {
		start = len(sessions)
	}
	if end > len(sessions) {
		end = len(sessions)
	}

	response := models.PaginatedResponse{
		Items: sessions[start:end],
		Pagination: models.Pagination{
			CurrentPage:  page,
			TotalPages:   totalPages,
			TotalItems:   totalItems,
			ItemsPerPage: limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// CreateGroup handles POST /api/groups
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.groupService.CreateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// UpdateGroup handles PUT /api/groups/:id
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group.ID = id
	if err := h.groupService.UpdateGroup(&group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

// DeleteGroup handles DELETE /api/groups/:id
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	if err := h.groupService.DeleteGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
