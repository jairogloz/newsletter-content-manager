package content_item

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Get handler")
}
