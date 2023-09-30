package model

import "time"

type LinkMediaModel struct {
	ID        int       `db:"id"`
	LinkID    int       `db:"link_id"`
	MediaPath string    `db:"media_path"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
