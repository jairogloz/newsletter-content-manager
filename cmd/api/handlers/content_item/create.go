package content_item

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/newsletter-content-manager/internal/domain"
	"net/http"
)

// Create is the handler for the POST /content-item endpoint.
func (h *Handler) Create(c *gin.Context) {
	var params domain.ContentItemCreateParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := h.Service.Create(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ooops! something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, newItem)
}
