package link

import "time"

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
