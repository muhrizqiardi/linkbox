package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

type FolderEntity struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	UniqueName string    `db:"unique_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type LinkEntity struct {
	ID          int       `db:"id"`
	URL         string    `db:"url"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	UserID      int       `db:"user_id"`
	FolderID    int       `db:"folder_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UserEntity struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
