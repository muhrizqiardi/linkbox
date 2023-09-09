package common

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
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

func HandleCreateUser(lg *log.Logger, us user.Service, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var payload user.CreateUserDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := us.Create(payload)
		if err != nil {
			lg.Println("failed to create user:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := as.LogIn(auth.LogInDTO{Username: user.Username, Password: user.Password})
		if err != nil {
			lg.Println("failed to log in:", err)
			http.Error(w, "Failed to log in. Account creation was success, so you can try logging in manually.", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:   "token",
			Value:  token,
			MaxAge: 7 * 24 * 60 * 60,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func HandleLogInPage(lg *log.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.LogInPage(w, templates.LogInPageData{}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleAuthLogIn(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var payload auth.LogInDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := as.LogIn(payload)
		if err != nil {
			lg.Println("failed to log in:", err)
			http.Error(w, "Failed to log in.", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:   "token",
			Value:  token,
			MaxAge: 7 * 24 * 60 * 60,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
