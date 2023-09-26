package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
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
}

func NewAuthHandler(lg *log.Logger, as service.AuthService, us service.UserService) *authHandler {
	return &authHandler{lg, as, us}
}

func (ah *authHandler) HandleAuthLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ah.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload request.LogInRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		ah.lg.Println("failed to decode form body into a struct:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := ah.as.LogIn(payload)
	if err != nil {
		ah.lg.Println("failed to log in:", err)
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

func (ah *authHandler) HandleCreateUserAndLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ah.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload request.CreateUserRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		ah.lg.Println("failed to decode form body into a struct:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := ah.us.Create(payload)
	if err != nil {
		ah.lg.Println("failed to create user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := ah.as.LogIn(request.LogInRequest{Username: user.Username, Password: payload.Password})
	if err != nil {
		ah.lg.Println("failed to log in:", err)
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
