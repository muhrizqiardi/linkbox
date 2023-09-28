package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/lib/pq"
	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/template"
	"github.com/muhrizqiardi/linkbox/internal/service"
)

type AuthHandler interface {
	HandleAuthLogIn(w http.ResponseWriter, r *http.Request)
	HandleCreateUserAndLogIn(w http.ResponseWriter, r *http.Request)
	HandleLogOut(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	lg *log.Logger
	as service.AuthService
	us service.UserService
	tx template.Executor
}

func NewAuthHandler(lg *log.Logger, as service.AuthService, us service.UserService, tx template.Executor) *authHandler {
	return &authHandler{lg, as, us, tx}
}

func (ah *authHandler) HandleAuthLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ah.lg.Println("failed to parse form body:", err)
		ah.tx.LogInPage(w, entities.LogInPageData{
			Errors: []string{
				constant.ErrLogInUser.Error(),
			},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
		return
	}
	var payload request.LogInRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		ah.lg.Println("failed to decode form body into a struct:", err)
		ah.tx.LogInPage(w, entities.LogInPageData{
			Errors: []string{
				constant.ErrLogInUser.Error(),
			},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
		return
	}

	token, err := ah.as.LogIn(payload)
	if err != nil {
		ah.lg.Println("failed to log in:", err)
		ah.tx.LogInPage(w, entities.LogInPageData{
			Errors: []string{
				constant.ErrLogInUser.Error(),
			},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
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

func (ah *authHandler) HandleCreateUserAndLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ah.lg.Println("failed to parse form body:", err)
		ah.tx.RegisterPage(w, entities.RegisterPageData{
			Errors: []string{constant.ErrRegisterUser.Error()},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
		return
	}
	var payload request.CreateUserRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		ah.lg.Println("failed to decode form body into a struct:", err)
		ah.tx.RegisterPage(w, entities.RegisterPageData{
			Errors: []string{constant.ErrRegisterUser.Error()},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
		return
	}
	user, err := ah.us.Create(payload)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch string(err.Code) {
			case "23505":
				ah.tx.RegisterPage(w, entities.RegisterPageData{
					Errors: []string{constant.ErrDuplicateUsername.Error()},
					PageMetaData: entities.PageMetaData{
						Title:       "Register Account - Linkbox",
						Description: "Register a new account on Linkbox",
						ImageURL:    "",
					},
				})
				return
			}
		}

		ah.lg.Println("failed to create user:", err)
		ah.tx.RegisterPage(w, entities.RegisterPageData{
			Errors: []string{constant.ErrRegisterUser.Error()},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
		return
	}
	token, err := ah.as.LogIn(request.LogInRequest{Username: user.Username, Password: payload.Password})
	if err != nil {
		ah.lg.Println("failed to log in after creating user:", err)
		ah.tx.RegisterPage(w, entities.RegisterPageData{
			Errors: []string{
				"Failed to log in. Account creation was success, so you can try logging in manually.",
			},
			PageMetaData: entities.PageMetaData{
				Title:       "Register Account - Linkbox",
				Description: "Register a new account on Linkbox",
				ImageURL:    "",
			},
		})
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

func (h *authHandler) HandleLogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/log-in", http.StatusSeeOther)
	return
}
