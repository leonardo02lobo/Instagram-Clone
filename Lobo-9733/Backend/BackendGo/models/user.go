package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserID       int            `json:"user_id" db:"user_id"`
	Username     string         `json:"username" db:"username"`
	Email        string         `json:"email" db:"email"`
	PasswordHash string         `json:"PasswordHash" db:"password_hash"`
	Bio          sql.NullString `json:"bio,omitempty" db:"bio"`
	ProfilePic   sql.NullString `json:"profile_pic,omitempty" db:"profile_pic"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
}
