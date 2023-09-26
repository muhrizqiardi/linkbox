package setup

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/jmoiron/sqlx"
)

func SetupRedisearch(lg *log.Logger, db *sqlx.DB) *redisearch.Client {
	connstring := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rsc := redisearch.NewClient(connstring, os.Getenv("REDIS_INDEX_NAME"))
	sch := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("title")).
		AddField(redisearch.NewTextField("description"))
	rsc.Drop()

	if err := rsc.CreateIndex(sch); err != nil {
		lg.Println("failed to create Redis index:", err)
	}

	q := `
		select id, url, title, description, user_id, folder_id, created_at, updated_at
			from links;
	`
	var ll []struct {
		ID          int       `db:"id"`
		UserID      int       `db:"user_id"`
		FolderID    int       `db:"folder_id"`
		URL         string    `db:"url"`
		Title       string    `db:"title"`
		Description string    `db:"description"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
	}
	if err := db.Select(&ll, q); err != nil {
		lg.Println("failed to query links:", err)
	}

	docs := make(redisearch.DocumentList, 0, len(ll))
	for i, e := range ll {
		doc := redisearch.NewDocument(fmt.Sprintf("doc%d", i), 1.0)
		doc.
			Set("id", e.ID).
			Set("user_id", e.UserID).
			Set("url", e.URL).
			Set("folder_id", e.FolderID).
			Set("title", e.Title).
			Set("description", e.Description).
			Set("created_at", e.CreatedAt.Unix()).
			Set("updated_at", e.UpdatedAt.Unix())
		docs = append(docs, doc)
	}

	if err := rsc.IndexOptions(redisearch.DefaultIndexingOptions, docs...); err != nil {
		lg.Println("failed to index in Redis:", err)
	}

	return rsc
}
