package content_item

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Update(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Update handler")
}
