package common

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func Route(lg *log.Logger, us user.Service, as auth.Service, fs folder.Service) chi.Router {
	r := chi.NewRouter()

	r.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	r.Handle("/node_modules/*", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))

	r.Get("/register", HandleRegisterPage(lg, as))
	r.Post("/register", HandleCreateUser(lg, us, as))

	r.Get("/log-in", HandleLogInPage(lg, as))
	r.Post("/log-in", HandleAuthLogIn(lg, as))

	// Needs Authentication
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(lg, as, us))

		r.Get("/", HandleIndexPage(lg, fs))
	})

	return r
}
