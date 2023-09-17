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
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/link"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/page"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/route"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/templates"
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

	db, err := sqlx.Connect("postgres", connstring)
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

	t, err := templates.NewTemplates()
	if err != nil {
		lg.Println("failed to instantiate templates:", err)
	}

	ur := user.NewRepository(db)
	fr := folder.NewRepository(db)
	lr := link.NewRepository(db)
	ls := link.NewService(lr)
	fs := folder.NewService(fr)
	us := user.NewService(ur, fs)
	as := auth.NewService(us, os.Getenv("SECRET"))
	ah := auth.NewHandler(lg, as, us)
	am := auth.NewMiddleware(lg, as, us)
	lh := link.NewHandler(lg, ls, t)
	fh := folder.NewHandler(lg, fs)

	ph := page.NewHandler(lg, fs, ls, as, t)
	r := route.Route(lg, ph, ah, am, lh, fh)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	lg.Fatalln(http.ListenAndServe(addr, r))
}
