package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConfig_NewFromMap(t *testing.T) {
	input := map[string]string{
		"PORT":                   "8080",
		"POSTGRES_USER":          "myuser",
		"POSTGRES_PASSWORD":      "testpassword",
		"POSTGRES_DB":            "mydb",
		"DB_HOST":                "localhost",
		"DB_PORT":                "5432",
		"DATABASE_URL":           "postgres://myuser:mypassword@localhost:5432/mydb",
		"TEST_POSTGRES_USER":     "testuser",
		"TEST_POSTGRES_PASSWORD": "testpassword",
		"TEST_POSTGRES_DB":       "testdb",
		"TEST_DB_HOST":           "localhost",
		"TEST_DB_PORT":           "5432",
		"REDIS_HOST":             "redis-server",
		"REDIS_PORT":             "6379",
		"REDIS_INDEX_NAME":       "myindex",
		"SECRET":                 "mysecret",
	}

	exp := &Config{
		Port:                 8080,
		PostgresUser:         "myuser",
		PostgresPassword:     "testpassword",
		PostgresDB:           "mydb",
		DBHost:               "localhost",
		DBPort:               "5432",
		DatabaseURL:          "postgres://myuser:mypassword@localhost:5432/mydb",
		TestPostgresUser:     "testuser",
		TestPostgresPassword: "testpassword",
		TestPostgresDB:       "testdb",
		TestDBHost:           "localhost",
		TestDBPort:           "5432",
		RedisHost:            "redis-server",
		RedisPort:            "6379",
		RedisIndexName:       "myindex",
		Secret:               "mysecret",
	}
	got, err := NewFromMap(input)
	if err != nil {
		t.Error("exp nil; got", err)
	}
	diff := cmp.Diff(exp, got)
	if diff != "" {
		t.Errorf("exp \"\"; got: \n %s", diff)
	}
}
