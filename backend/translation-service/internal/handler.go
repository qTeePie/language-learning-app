package translate

import (
	"fmt"
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
	text := c.Query("text")
	pos := c.Query("pos")

	fmt.Println("Stepped into handler...")

	translation, err := h.service.GetTranslation(sourceLang, targetLang, text, pos)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"translation": translation})
}
