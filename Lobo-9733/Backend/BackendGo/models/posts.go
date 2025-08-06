package models

import "time"

type Posts struct {
	PostID    int       `json:"post_id"`
	UserId    int       `json:"user_id"`
	Caption   string    `json:"caption"`
	MediaUrl  string    `json:"media_url"`
	MediaType string    `json:"media_type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
