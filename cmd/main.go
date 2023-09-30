package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	files "github.com/muhrizqiardi/linkbox"
	"github.com/muhrizqiardi/linkbox/internal/config"
	"github.com/muhrizqiardi/linkbox/internal/db"
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

func main() {
	lg := log.New(os.Stdout, "linkbox | ", log.LstdFlags)

	envMap, err := godotenv.Read(".env")
	if err != nil {
		lg.Fatalln("failed to retrieve environment variable:", err)
	}
	cfg, err := config.NewFromMap(envMap)
	if err != nil {
		lg.Fatalln("failed to assign environment variables:", err)
	}

	db, err := db.ConnectPostgresWithMigration(
		lg,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.DBHost,
		"disable", // TODO: change from disable
	)
	if err != nil {
		lg.Fatalln("failed to connect to database:", err)
	}
	defer db.Close()
	lg.Println("successfully connected to database")

	dbx := sqlx.NewDb(db, "postgres")
	rsc := setup.SetupRedisearch(lg, dbx, cfg.RedisHost, cfg.RedisPort, cfg.RedisIndexName)

	t, err := template.NewExecutor()
	if err != nil {
		lg.Println("failed to instantiate templates:", err)
	}

	lr := repository.NewLinkRepository(dbx, rsc)
	lmr := repository.NewLinkMediaRepository(dbx)
	fr := repository.NewFolderRepository(dbx)
	ur := repository.NewUserRepository(dbx)
	ls := service.NewLinkService(lr, lmr)
	fs := service.NewFolderService(fr)
	us := service.NewUserService(ur)
	as := service.NewAuthService(ur, cfg.Secret)
	am := middleware.NewAuthMiddleware(lg, as, us)
	ah := handler.NewAuthHandler(lg, as, us, t)
	lh := handler.NewLinkHandler(lg, ls, t, fs)
	fh := handler.NewFolderHandler(lg, fs)
	ph := handler.NewPageHandler(lg, fs, ls, as, t)
	r := route.DefineRoute(lg, ph, ah, am, lh, fh, files.DistFS)

	addr := fmt.Sprintf(":%d", cfg.Port)
	lg.Fatalln(http.ListenAndServe(addr, r))
}
