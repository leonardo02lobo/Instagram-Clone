package models

import "time"

type Like struct {
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
