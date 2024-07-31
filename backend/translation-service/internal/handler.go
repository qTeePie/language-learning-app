package translate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TranslationHandler handles translation requests
type Handler struct {
	service *Service
}

// NewTranslationHandler creates a new TranslationHandler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// HandleTranslation handles the translation request
func (h *Handler) HandleTranslation(c *gin.Context) {
	sourceLang := c.Query("source_lang")
	targetLang := c.Query("target_lang")
	word := c.Query("word")
	pos := c.Query("pos")

	translation, err := h.service.GetTranslation(sourceLang, targetLang, word, pos)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Translation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"translation": translation})
}
