package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/muhrizqiardi/linkbox/internal/config"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/handler"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/middleware"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/route"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/template"
	"github.com/muhrizqiardi/linkbox/internal/repository"
	"github.com/muhrizqiardi/linkbox/internal/service"
	"github.com/muhrizqiardi/linkbox/internal/util/setup"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func setupDB(cfg *config.Config) (*sqlx.DB, error) {
	connstring := fmt.Sprintf(
		"user='%s' password='%s' dbname='%s' host='%s' sslmode='disable'",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.DBHost,
	)

	db, err := sqlx.Connect("postgres", connstring)
	if err != nil {
		return &sqlx.DB{}, err
	}

	return db, nil
}

func main() {
	lg := log.New(os.Stdout, "linkbox | ", log.LstdFlags)

	envMap, err := godotenv.Read(".env")
	if err != nil {
		defer lg.Println("tips: make sure .env has been created")
		lg.Fatalln("failed to retrieve environment variable:", err)
	}
	cfg, err := config.NewFromMap(envMap)
	if err != nil {
		defer lg.Println("tips: make sure .env file is in the correct format, check .env.example file")
		lg.Fatalln("failed to assign environment variables:", err)
	}

	db, err := setupDB(cfg)
	if err != nil {
		defer lg.Println("tips: make sure the database has been started")
		lg.Fatalln("failed to connect to database:", err)
	}
	defer lg.Println("database connection closed")
	defer db.Close()
	lg.Println("successfully connected to database")

	rsc := setup.SetupRedisearch(lg, db)

	t, err := template.NewExecutor()
	if err != nil {
		lg.Println("failed to instantiate templates:", err)
	}

	lr := repository.NewLinkRepository(db, rsc)
	fr := repository.NewFolderRepository(db)
	ur := repository.NewUserRepository(db)
	ls := service.NewLinkService(lr)
	fs := service.NewFolderService(fr)
	us := service.NewUserService(ur, fs)
	as := service.NewAuthService(us, cfg.Secret)
	am := middleware.NewAuthMiddleware(lg, as, us)
	ah := handler.NewAuthHandler(lg, as, us)
	lh := handler.NewLinkHandler(lg, ls, t)
	fh := handler.NewFolderHandler(lg, fs)
	ph := handler.NewPageHandler(lg, fs, ls, as, t)
	r := route.DefineRoute(lg, ph, ah, am, lh, fh)

	addr := fmt.Sprintf(":%d", cfg.Port)
	lg.Fatalln(http.ListenAndServe(addr, r))
}
