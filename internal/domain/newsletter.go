package domain

import "time"

// Newsletter represents an instance of a newsletter that has been published. It holds references to all the items that are part of the newsletter.
type Newsletter struct {
	ID          interface{}   `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	ItemIDs     []interface{} `json:"item_ids" bson:"item_ids"`
	PublishedAt *time.Time    `json:"published_at" bson:"published_at"`
}
