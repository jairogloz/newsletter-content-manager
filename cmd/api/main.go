package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/newsletter-content-manager/cmd/api/core"
	"github.com/jairogloz/newsletter-content-manager/cmd/api/handlers/content_item"
	contentItemService "github.com/jairogloz/newsletter-content-manager/internal/services/content_item"
	"log"
)

func main() {

	contentItemSrv := contentItemService.NewService()

	server := core.Server{
		GinEngine:          gin.Default(),
		ContentItemHandler: content_item.NewHandler(contentItemSrv),
	}

	server.GinEngine.POST("/content-item", server.ContentItemHandler.Create)
	server.GinEngine.GET("/content-item", server.ContentItemHandler.Get)
	server.GinEngine.PUT("/content-item", server.ContentItemHandler.Update)
	server.GinEngine.DELETE("/content-item", server.ContentItemHandler.Delete)

	log.Fatalln(server.GinEngine.Run(":8001"))
}
