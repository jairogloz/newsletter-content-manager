package content_item

import (
	"github.com/jairogloz/newsletter-content-manager/internal/ports/service"
)

// Handler is the struct that holds the handlers for the content item endpoints.
// It contains the required components to handle the requests and responses for the content item domain.
type Handler struct {
	Service service.ContentItemService
}

// NewHandler creates a new instance of the content item handler.
func NewHandler(service service.ContentItemService) *Handler {
	return &Handler{
		Service: service,
	}
}
