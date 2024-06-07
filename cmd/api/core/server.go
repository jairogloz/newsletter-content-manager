package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/newsletter-content-manager/cmd/api/handlers/content_item"
)

// Server represents the core server and holds the required components to expose
// the application services through http endpoints using the gin framework.
type Server struct {
	GinEngine          *gin.Engine
	ContentItemHandler *content_item.Handler
}
