package api

import (
	"net/http"
	"strconv"

	"github.com/erans/lang-portal/internal/models"
	"github.com/erans/lang-portal/internal/service"
	"github.com/gin-gonic/gin"
)

// WordHandler handles HTTP requests for words
type WordHandler struct {
	wordService *service.WordService
}

// NewWordHandler creates a new WordHandler
func NewWordHandler(wordService *service.WordService) *WordHandler {
	return &WordHandler{wordService: wordService}
}

// RegisterRoutes registers the word routes
func (h *WordHandler) RegisterRoutes(r *gin.Engine) {
	words := r.Group("/api/words")
	{
		words.GET("", h.ListWords)
		words.GET("/:id", h.GetWord)
		words.POST("", h.CreateWord)
		words.PUT("/:id", h.UpdateWord)
		words.DELETE("/:id", h.DeleteWord)
	}
}

// ListWords handles GET /api/words
func (h *WordHandler) ListWords(c *gin.Context) {
	// Get pagination parameters with default 100 items per page as per spec
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 100 // Fixed as per spec
	offset := (page - 1) * limit

	// Get words from service
	result, err := h.wordService.ListWords(offset, limit)
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

// GetWord handles GET /api/words/:id
func (h *WordHandler) GetWord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	word, err := h.wordService.GetWord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}

// CreateWord handles POST /api/words
func (h *WordHandler) CreateWord(c *gin.Context) {
	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.wordService.CreateWord(&word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, word)
}

// UpdateWord handles PUT /api/words/:id
func (h *WordHandler) UpdateWord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word.ID = id
	if err := h.wordService.UpdateWord(&word); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}

// DeleteWord handles DELETE /api/words/:id
func (h *WordHandler) DeleteWord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.wordService.DeleteWord(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
