package response

import (
	"time"
)

type LinkWithMediaResponseMedia struct {
	MediaPath *string `db:"media_path"`
}

type LinkWithMediaResponse struct {
	ID          int    `db:"id"`
	URL         string `db:"url"`
	Title       string `db:"title"`
	Description string `db:"description"`
	UserID      int    `db:"user_id"`
	FolderID    int    `db:"folder_id"`
	Media       []LinkWithMediaResponseMedia
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
