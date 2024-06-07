package content_item

import (
	"github.com/jairogloz/newsletter-content-manager/internal/ports/service"
)

// Service implements service.ContentItemService interface, and holds the required
// components to handle the business logic for the content item domain.
type Service struct {
}

// NewService creates a new instance of the content item service.
func NewService() service.ContentItemService {
	return &Service{}
}
