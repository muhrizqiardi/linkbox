package common

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/templates"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func HandleIndexPage(lg *log.Logger, fs folder.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uCtx := r.Context().Value("user")
		foundUser, ok := uCtx.(user.UserEntity)
		if !ok {
			lg.Println("failed to get user data passed from middleware")
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		folders, err := fs.GetMany(foundUser.ID, folder.GetManyFoldersDTO{
			Sort:    folder.GetManyFoldersSortDescending,
			OrderBy: folder.GetManyFoldersOrderByUpdatedAt,
			Limit:   20,
			Offset:  0,
		})
		if err != nil {
			lg.Println("failed to find folders related to user", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		if err := templates.IndexPage(w, templates.IndexPageData{
			User:    foundUser,
			Folders: folders,
		}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleRegisterPage(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		existingCookie, err := r.Cookie("token")
		if err == nil {
			_, newToken, err := as.CheckIsValid(existingCookie.Value)
			if err == nil {
				lg.Println("user already authenticated, redirecting")
				http.SetCookie(w, &http.Cookie{
					Name:   "token",
					Value:  newToken,
					MaxAge: 8 * 24 * 60 * 60,
				})
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		if err := templates.RegisterPage(w, templates.RegisterPageData{}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
		return
	}
}

func HandleLogInPage(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		existingCookie, err := r.Cookie("token")
		if err == nil {
			_, newToken, err := as.CheckIsValid(existingCookie.Value)
			if err == nil {
				lg.Println("user already authenticated, redirecting")
				http.SetCookie(w, &http.Cookie{
					Name:   "token",
					Value:  newToken,
					MaxAge: 8 * 24 * 60 * 60,
				})
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

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
		return
	}
}
