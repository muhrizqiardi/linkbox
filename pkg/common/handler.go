package common

import (
	"log"
	"net/http"

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
		// user, err := us.Create()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
