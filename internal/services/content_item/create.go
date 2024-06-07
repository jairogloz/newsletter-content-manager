package content_item

import (
	"context"
	"github.com/jairogloz/newsletter-content-manager/internal/domain"
	"time"
)

// Create creates a new content item and stores it in the database.
func (s Service) Create(ctx context.Context, params domain.ContentItemCreateParams) (*domain.ContentItem, error) {
	now := time.Now().UTC()
	contentItem := &domain.ContentItem{
		ID:          "123",
		Title:       params.Title,
		Content:     params.Content,
		Type:        params.Type,
		Category:    params.Category,
		Tags:        params.Tags,
		Author:      params.Author,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		PublishedIn: nil,
	}
	return contentItem, nil
}
