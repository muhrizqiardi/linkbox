package integration

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/muhrizqiardi/linkbox/internal/db"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/query"
)

func setupTests(t *testing.T) (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		t.Error("failed to retrieve environment variable:", err)
		return &sqlx.DB{}, err
	}

	postgres, err := db.ConnectPostgresWithMigration(
		log.Default(),
		os.Getenv("TEST_POSTGRES_USER"),
		os.Getenv("TEST_POSTGRES_PASSWORD"),
		os.Getenv("TEST_POSTGRES_DB"),
		os.Getenv("TEST_DB_HOST"),
		"disable",
	)
	if err != nil {
		t.Error("failed to connect to database:", err)
		return &sqlx.DB{}, err
	}

	return sqlx.NewDb(postgres, "postgres"), nil
}

func TestQueryCreateUser(t *testing.T) {
	postgres, err := setupTests(t)
	if err != nil {
		t.Fatal("failed to setup test:", err)
	}

	mockUsername := "johndoe"
	mockPassword := "topsecret"

	var u model.UserModel
	if err := postgres.Get(&u, query.QueryCreateUserWithDefaultFolder, mockUsername, mockPassword); err != nil {
		t.Error("exp nil, got error:", err)
	}

	var defaultFolder model.FolderModel
	if err := postgres.Get(
		&defaultFolder,
		`select * from folders where user_id = $1 and unique_name = 'default';`,
		u.ID,
	); err != nil {
		t.Error("exp nil, got error:", err)
	}

	{
		exp := "default"
		got := defaultFolder.UniqueName
		if exp != got {
			t.Errorf("exp %s; got %s", exp, got)
		}
	}

	{
		exp := u.ID
		got := defaultFolder.UserID
		if exp != got {
			t.Errorf("exp %d; got %d", exp, got)
		}
	}
}
