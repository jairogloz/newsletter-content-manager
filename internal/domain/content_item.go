package domain

import (
	"time"
)

// ContentItem represents a piece of content that can be published in a newsletter or any other content feed.
type ContentItem struct {
	ID          interface{}  `json:"id" bson:"_id"`
	Title       string       `json:"title" bson:"title"`
	Content     string       `json:"content" bson:"content"`
	Type        string       `json:"type" bson:"type"`
	Category    string       `json:"category" bson:"category"`
	Tags        []string     `json:"tags" bson:"tags"`
	Author      string       `json:"author" bson:"author"`
	CreatedAt   *time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at" bson:"updated_at"`
	PublishedIn *PublishedIn `json:"published_in" bson:"published_in"`
}

// ContentItemCreateParams represents the parameters needed to create a new content item.
type ContentItemCreateParams struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Type     string   `json:"type"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Author   string   `json:"author"`
}

// PublishedIn holds a list of means where the content item can be published in. It helps
// to keep track of the different content feeds where the item has been published. For instance,
// a content item can be published in a newsletter already but not in a blogpost.
type PublishedIn struct {
	Newsletter bool `json:"newsletter" bson:"newsletter"`
	Youtube    bool `json:"youtube" bson:"youtube"`
}
