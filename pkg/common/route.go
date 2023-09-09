package common

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func Route(lg *log.Logger, us user.Service) chi.Router {
	r := chi.NewRouter()

	r.Get("/register", HandleRegisterPage(lg))
	r.Post("/register", HandleCreateUser(lg, us))

	return r
}
