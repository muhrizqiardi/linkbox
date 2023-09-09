package common

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/templates"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func HandleRegisterPage(lg *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.RegisterPage(w, templates.RegisterPageData{}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func HandleCreateUser(lg *log.Logger, us user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var payload user.CreateUserDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if _, err := us.Create(payload); err != nil {
			lg.Println("failed to create user:", err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
