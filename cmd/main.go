package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func setupDB() (*sqlx.DB, error) {
	connstring := fmt.Sprintf(
		"user='%s' password='%s' dbname='%s' host='%s' sslmode='disable'",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_HOST"),
	)

	db, err := sqlx.Open("postgres", connstring)
	if err != nil {
		return &sqlx.DB{}, err
	}

	return db, nil
}

func main() {
	lg := log.New(os.Stdout, "linkbox | ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		lg.Fatalln("failed to retrieve environment variables:", err)
	}

	db, err := setupDB()
	if err != nil {
		lg.Fatalln("failed to connect to database:", err)
	}
	defer db.Close()
	lg.Println("successfully connected to database")

	ur := user.NewRepository(db)
	fr := folder.NewRepository(db)
	fs := folder.NewService(fr)
	us := user.NewService(ur, fs)
	as := auth.NewService(us, os.Getenv("SECRET"))
	r := common.Route(lg, us, as)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	lg.Fatalln(http.ListenAndServe(addr, r))
}
