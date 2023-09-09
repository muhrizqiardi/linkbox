package folder

import "time"

type FolderEntity struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	UniqueName string    `db:"unique_name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
