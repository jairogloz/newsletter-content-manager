package content_item

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Delete(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Delete handler!!")
}
