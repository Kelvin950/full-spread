package domain

import "time"

type Post struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Type        PostType  `json:"type"`
	Published   bool      `json:"published"`
	CreatorID   uint      `json:"creator_id"`
	Content     []Content `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Content struct {
	ID              uint      `json:"id"`
	MimeType        string    `json:"mime_type"`
	LocationUrl     string    `json:"location_url"`
	PostID          uint      `json:"post_id"`
	ManifestFileUrl *string   `json:"manifest_file_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
