package service

import (
	"context"
	"github.com/jairogloz/newsletter-content-manager/internal/domain"
)

// ContentItemService exposes the services provided by this application to
// handle content items operations.
type ContentItemService interface {
	Create(ctx context.Context, params domain.ContentItemCreateParams) (*domain.ContentItem, error)
}
